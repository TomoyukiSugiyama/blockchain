package blockchain

import (
	"blockchain/block"
	"fmt"
	"time"
)

type Blockchain struct {
	Blocks map[string]*block.Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.createGenesisBlock()

	// Test adding a block
	b := block.Block{
		Index:     1,
		Timestamp: time.Now().String(),
		Nonce:     0,
		Data:      "Hi there!",
		PrevHash:  "0000",
	}
	bc.addBlock(&b, "000011")
	return bc
}

func (bc *Blockchain) createGenesisBlock() {
	b := block.Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Nonce:     1,
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	b.Hash = block.CalculateHash(&b)
	fmt.Println("Genesis Block Hash: ", b.Hash)
	bc.Blocks = make(map[string]*block.Block)
	bc.Blocks["0000"] = &b
	block.Add(bc.Blocks["0000"])
}

const difficulty = 4

func (bc *Blockchain) addBlock(b *block.Block, h string) {
	if !checkHash(h, difficulty) {
		// Reject block
		return
	}

	if !isExistPreviousBlock(bc.Blocks, b.PrevHash) {
		// Reject block
		return
	}

	bc.Blocks[h] = b
	block.Add(bc.Blocks[h])
}

func checkHash(hash string, difficulty int) bool {
	for i := 0; i < difficulty; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

func isExistPreviousBlock(blocks map[string]*block.Block, prevHash string) bool {
	if _, ok := blocks[prevHash]; ok {
		return true
	}

	return false
}
