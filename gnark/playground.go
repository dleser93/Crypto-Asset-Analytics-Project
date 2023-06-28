// Welcome to the gnark playground!
package main

import (
	"github.com/consensys/gnark/frontend"
	// "github.com/consensys/gnark/std/hash/mimc"
)

const Max = (1 << 16) - 1
const MaxLength = 2

// gnark is a zk-SNARK library written in Go. Circuits are regular structs.
// The inputs must be of type frontend.Variable and make up the witness.
// The witness has a
//   - secret part --> known to the prover only
//   - public part --> known to the prover and the verifier
type Circuit struct {

	ClientIndex   frontend.Variable `gnark:",public"` // known to all
	ClientId      frontend.Variable `gnark:",public"` // known to all
	ClientBalance frontend.Variable `gnark:",public"` // known to all

	Hashes [MaxLength]frontend.Variable // known to the prover (!= verifier) only
	Sums [MaxLength]frontend.Variable // known to the prover (!= verifier) only

	RootHash frontend.Variable `gnark:",public"` // known to all (in this case really ALL)
	RootSum  frontend.Variable `gnark:",public"` // known to all (in this case really ALL)
}

// Define declares the circuit logic. The compiler then produces a list of constraints
// which must be satisfied (valid witness) in order to create a valid zk-SNARK
// This circuit proves knowledge of a pre-image such that hash(secret) == hash
func (circuit *Circuit) Define(api frontend.API) error {


	api.AssertIsLessOrEqual(0, circuit.ClientBalance)
	api.AssertIsLessOrEqual(circuit.ClientBalance, Max)

	helper := api.ToBinary(circuit.ClientIndex, MaxLength)

	// hash function
	// mimc, _ := mimc.NewMiMC(api)

	// hash the secret
	// mimc.Write(circuit.ClientId)

	// ensure hashes match
	// nodeHash := mimc.Sum()

	nodeHash := circuit.ClientId
	nodeSum  := circuit.ClientBalance

	for i := 0; i < MaxLength; i++ {
		h1 := api.Select(helper[i], circuit.Hashes[i], nodeHash)
		h2 := api.Select(helper[i], nodeHash, circuit.Hashes[i])

		// mimc.Reset()
		// mimc.Write(h1)
		// mimc.Write(h2)
		// nodeHash = mimc.Sum()

		nodeHash = api.Mul(h1, h2)

		api.AssertIsLessOrEqual(0, circuit.Sums[i])
		api.AssertIsLessOrEqual(circuit.Sums[i], Max)

		nodeSum = api.Add(nodeSum, circuit.Sums[i])

		api.AssertIsLessOrEqual(0, nodeSum)
		api.AssertIsLessOrEqual(nodeSum, Max)
	}

	// Compare our calculated Merkle root and sum to the desired Merkle root and sum.
	api.AssertIsEqual(nodeHash, circuit.RootHash)
	api.AssertIsEqual(nodeSum, circuit.RootSum)

	return nil
}

-- witness.json --
{
        "ClientIndex": "0",
        "ClientId": "161",
        "ClientBalance": "100",

        "Hashes": ["178", "41340"],
        "Sums": ["200", "700"],

        "RootHash": "1184721720",
        "RootSum": "1000"
}
