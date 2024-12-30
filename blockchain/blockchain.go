package blockchain

import (
	"blockchain/block"
	"fmt"
	"time"
)

type Blockchain struct {
	Blocks       map[string]*block.Block
	currentBlock *block.Block
}

const tryLimit = 1000000
const difficulty = 4

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.createGenesisBlock()

	// Test adding a block
	bc.MineBlock("First Block")
	bc.MineBlock("Second Block")
	bc.MineBlock("Third Block")
	bc.MineBlock("Fourth Block")
	return bc
}

func (bc *Blockchain) createGenesisBlock() {
	genesisBlock := block.Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Nonce:     1,
		Data:      "Genesis Block",
		PrevHash:  "",
	}

	for i := 0; i < tryLimit; i++ {
		genesisBlock.Nonce = i
		genesisBlock.Hash = genesisBlock.CalculateHash()
		if checkHash(genesisBlock.Hash, difficulty) {
			break
		}
	}

	fmt.Println("Genesis Block Hash: ", genesisBlock.Hash)
	bc.Blocks = make(map[string]*block.Block)
	bc.currentBlock = &genesisBlock
	bc.Blocks[genesisBlock.Hash] = &genesisBlock
	bc.Blocks[genesisBlock.Hash].Add()
}

func (bc *Blockchain) addBlock(b *block.Block, h string) {
	// if !checkHash(h, difficulty) {
	// 	// Reject block
	// 	return
	// }

	if !isExistPreviousBlock(bc.Blocks, b.PrevHash) {
		// Reject block
		return
	}

	bc.Blocks[h] = b
	bc.Blocks[h].Add()
}

func (bc *Blockchain) MineBlock(message string) {
	newBlock := bc.currentBlock.GenerateBlock()
	newBlock.Data = message
	for i := 0; i < tryLimit; i++ {
		newBlock.Nonce = i
		newBlock.Hash = newBlock.CalculateHash()
		if checkHash(newBlock.Hash, difficulty) {
			break
		}
	}
	bc.addBlock(newBlock, newBlock.Hash)
	bc.currentBlock = newBlock
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
