package main

import (
	// "github.com/consensys/gnark-crypto/ecc/bls12-381/fr/mimc"
	// "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"fmt"
)

type Tuple struct {
	Identifier uint64
	Balance    uint64
}

type MerkleNode struct {
	LeftChild  *MerkleNode
	RightChild *MerkleNode
	// Hash       []byte
	Hash       uint64
	Sum        uint64
}

func main() {

	// only power-of-two number of elements!
	tuples := []Tuple{
		{Identifier: 0xa1, Balance: 100},
		{Identifier: 0xb2, Balance: 200},
		{Identifier: 0xc3, Balance: 300},
		{Identifier: 0xd4, Balance: 400},
	}

	clientIndex := uint64(1)
	merkleTreeDepth := 2

	merkleRoot := buildMerkleTree(tuples)
	// fmt.Printf("Merkle Root: %x Sum: %d Length: %d\n", merkleRoot.Hash, merkleRoot.Sum, len(merkleRoot.Hash))
	// fmt.Printf("Merkle Root: %d Sum: %d", merkleRoot.Hash, merkleRoot.Sum)

	id, bal, hashes, sums := GetMerkleProof(clientIndex, merkleTreeDepth, merkleRoot, tuples)

	// fmt.Println(id, bal, hashes, sums)
	fmt.Println("{")
	fmt.Printf("	\"ClientIndex\": \"%d\",\n", clientIndex)
    fmt.Printf("	\"ClientId\": \"%d\",\n", id)
    fmt.Printf("	\"ClientBalance\": \"%d\",\n", bal)
	fmt.Println("")
    fmt.Printf("	\"Hashes\": [\"%d\"%s],\n", hashes[0], printUint64s(hashes[1:]))
    fmt.Printf("	\"Sums\": [\"%d\"%s],\n", sums[0], printUint64s(sums[1:]))
    fmt.Println("")
    fmt.Printf("	\"RootHash\": \"%d\",\n", merkleRoot.Hash)
    fmt.Printf("	\"RootSum\": \"%d\"\n", merkleRoot.Sum)
    fmt.Println("}")
}

func buildMerkleTree(tuples []Tuple) *MerkleNode {
	nodes := make([]*MerkleNode, len(tuples))

	// Create leaf nodes
	for i, tuple := range tuples {

		// data := make([]byte, 32)
		// copyStringToArray(fmt.Sprintf("%d:%d", tuple.Identifier, tuple.Balance), data)
		// fmt.Println(data)
		// hash, _ := mimc.Sum(data)

		hash := tuple.Identifier
		// nodes[i] = &MerkleNode{Hash: hash[:], Sum: tuple.Balance}
		nodes[i] = &MerkleNode{Hash: hash, Sum: tuple.Balance}
	}

	// Build tree from leaf nodes
	for len(nodes) > 1 {
		var levelNodes []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			leftChild := nodes[i]
			rightChild := nodes[i+1]

			// concatData := append(leftChild.Hash, rightChild.Hash...)
			// // fmt.Printf("%x\n%x\n%x\n", leftChild.Hash, rightChild.Hash, concatData)
			// hash, _ := mimc.Sum(concatData)

			hash := leftChild.Hash * rightChild.Hash

			parent := &MerkleNode{
				LeftChild:  leftChild,
				RightChild: rightChild,
				// Hash:       hash[:],
				Hash:       hash,
				Sum:		leftChild.Sum + rightChild.Sum,
			}

			levelNodes = append(levelNodes, parent)
		}

		nodes = levelNodes
	}

	return nodes[0]
}

// func GetMerkleProof(index uint64, depth int, root *MerkleNode, tuples []Tuple) (id, balance uint64, hashes [][]byte, sums []uint64, helper []bool) {
func GetMerkleProof(index uint64, depth int, root *MerkleNode, tuples []Tuple) (id, balance uint64, hashes, sums []uint64) {

	id = tuples[index].Identifier
	balance = tuples[index].Balance
	helper := getBooleanArray(index, depth)

	node := root

	for i := 0; i < len(helper); i++ {

		// fmt.Printf("Here %x %d\n", node.Hash, node.Sum)

		partner := SelectNodes(helper[i], node.LeftChild, node.RightChild)
		// fmt.Printf("Take %x %d\n", partner.Hash, partner.Sum)

		node = SelectNodes(helper[i], node.RightChild, node.LeftChild)

		hashes = append(hashes, partner.Hash)
		sums = append(sums, partner.Sum)
	}

	// fmt.Printf("Here %x %d\n", node.Hash, node.Sum)

	reverse(hashes)
	reverse(sums)

	return
}


func getBooleanArray(number uint64, depth int) []bool {
	var boolArray []bool

	for number > 0 {
		bit := number & 1
		boolArray = append(boolArray, bit == 1)
		number >>= 1
	}

	// fill with false
	for len(boolArray) < depth {
		boolArray = append(boolArray, false)
	}

	// Reverse the boolean array to maintain the correct order
	reverseSlice(boolArray)

	return boolArray
}


func reverseSlice(slice []bool) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func reverse[T any] (slice []T) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func copyStringToArray(str string, array []byte) {
	copy(array[:], []byte(str))
}

func SelectUint64(sel bool, l uint64, r uint64) uint64 {
	if(sel) {
		return l
	} else {
		return r
	}
}

func SelectBytes(sel bool, l []byte, r []byte) []byte {
	if(sel) {
		return l
	} else {
		return r
	}
}

func SelectNodes(sel bool, l *MerkleNode, r *MerkleNode) *MerkleNode {
	if(sel) {
		return l
	} else {
		return r
	}
}

func Btoi(b bool) int {
    if b {
        return 1
    }
    return 0
 }

func printBytes(arr [][]byte) string {

	str := ""

	for _, el := range arr {
		/* code */
		str = str + fmt.Sprintf("\"0x%x\", ", el)
	}

    return str
}

func printBools(arr []bool) string {

	str := ""

	for _, el := range arr {
		/* code */
		str = str + fmt.Sprintf("%d, ", Btoi(el))
	}

    return str
}

func printUint64s(arr []uint64) string {

	str := ""

	for _, el := range arr {
		/* code */
		str = str + fmt.Sprintf(", \"%d\"", el)
	}

    return str
}
