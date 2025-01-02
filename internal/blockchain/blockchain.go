package blockchain

import (
	"blockchain/internal/account"
	"blockchain/internal/block"
	"blockchain/internal/state"
	"log"
	"time"
)

type Blockchain struct {
	State []*state.State
}

const tryLimit = 1000000
const difficulty = 4

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}

	return bc
}

func (bc *Blockchain) CreateGenesisBlock() {
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

	bc.State = append(bc.State, state.CreateNewState(nil, &genesisBlock))
	log.Println(bc.State[len(bc.State)-1].String())
}

func (bc *Blockchain) addBlock(b *block.Block, h string, accs map[string]*account.Account) {
	if !checkHash(h, difficulty) {
		// Reject block
		log.Fatalln("Block Rejected with hash: ", h)
		return
	}

	if !isExistPreviousBlock(bc.State[len(bc.State)-1].Block.Hash, b.PrevHash) {
		// Reject block
		log.Fatalln("Block Rejected")
		return
	}

	bc.State = append(bc.State, state.CreateNewState(accs, b))
}

func (bc *Blockchain) MineBlock(message string, trs []block.Transaction, accs map[string]*account.Account) {
	newBlock := bc.State[len(bc.State)-1].Block.GenerateBlock()
	newBlock.Data = message
	newBlock.Transactions = trs
	for i := 0; i < tryLimit; i++ {
		newBlock.Nonce = i
		newBlock.Hash = newBlock.CalculateHash()
		if checkHash(newBlock.Hash, difficulty) {
			break
		}
	}
	bc.addBlock(newBlock, newBlock.Hash, accs)
	for _, t := range trs {
		t.Run(accs)
	}
	log.Println(bc.State[len(bc.State)-1].String())
}

func checkHash(hash string, difficulty int) bool {
	for i := 0; i < difficulty; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

func isExistPreviousBlock(hash string, prevHash string) bool {
	return hash == prevHash
}
