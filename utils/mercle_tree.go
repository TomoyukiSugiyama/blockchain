package utils

import (
	"encoding/json"
)

type MerkleNode struct {
	Hash  []byte      `json:"hash"`
	Left  *MerkleNode `json:"left"`
	Right *MerkleNode `json:"right"`
}

func BuildMerkleTree(data [][]byte) *MerkleNode {
	var nodes []*MerkleNode

	// Create leaf nodes
	for _, datum := range data {
		node := &MerkleNode{Hash: CalculateHash(datum)}
		nodes = append(nodes, node)
	}

	// Build tree
	for len(nodes) > 1 {
		var level []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := &MerkleNode{Hash: nil}
			if i+1 < len(nodes) {
				right = nodes[i+1]
			}
			parent := &MerkleNode{
				Hash:  CalculateHash(append(left.Hash, right.Hash...)),
				Left:  left,
				Right: right,
			}
			level = append(level, parent)
		}
		nodes = level
	}

	return nodes[0]
}

func (mn *MerkleNode) ToJson() []byte {
	jsonData, err := json.Marshal(mn)
	if err != nil {
		panic(err)
	}
	return jsonData
}
