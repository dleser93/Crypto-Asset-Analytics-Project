# Crypto-Asset-Analytics-Project - (zk-)SNARKed Merkle Sum Tree

This implementation is based on the idea of only using the zk-SNARK framework on the Merkle proof to make up for the weakness of the Merkle Sum Tree of leaking (sums of) balances to other users (https://ethresear.ch/t/snarked-merkle-sum-tree-a-practical-proof-of-solvency-protocol-based-on-vitaliks-proposal/14405).

The framework used is gnark (https://github.com/ConsenSys/gnark), a fork of which is used by Binance for their proof-of-reserves.

For a proof of concept two pieces of code are provided, "offline.go" and "playground.go".

In "offline.go" a power of two number of clients can be specified in the "tuples" array.
The depth d of the resulting Merkle Sum Tree has to be provided as well in "merkleTreeDepth" (2^d == len(tuples)).
Finally the position of the client in the "tuples" array for whom the Merkle Sum Tree proof shall be generated has to be set in clientIndex as an uint64.

Running the code with "go run offline.go" produces the "witness.json" data required by the prover to construct the zk-SNARK proof.

In "playground.go" the lines after the "--witness.json--" have to be replaced by the output of "offline.go" and the whole content of the file then can be checked on gnark's playground (https://play.gnark.io/) by copy-pasting it there and clicking the "Run" button.
