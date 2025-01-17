package blockchain

import (
	"blockchain/internal/account"
	"blockchain/internal/block"
	"blockchain/internal/state"
	"blockchain/internal/transaction"
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

func (bc *Blockchain) CreateGenesisBlock(acc map[string]*account.Account) {
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

	bc.State = append(bc.State, state.CreateNewState(acc, &genesisBlock))
	log.Printf("Create Genesis Block: %s", bc.State[len(bc.State)-1].ToJson())
}

func (bc *Blockchain) AddBlock(b *block.Block, txs []*transaction.Transaction, accs map[string]*account.Account) {
	if !checkHash(b.Hash, difficulty) {
		// Reject block
		log.Fatalln("Block Rejected with hash: ", b.Hash)
		return
	}

	if !isExistPreviousBlock(bc.State[len(bc.State)-1].Block.Hash, b.PrevHash) {
		// Reject block
		log.Fatalln("Block Rejected")
		return
	}

	bc.State = append(bc.State, state.CreateNewState(accs, b))
	for _, t := range txs {
		t.Run(accs)
	}
	log.Printf("Add Block: %s", bc.State[len(bc.State)-1].ToJson())
}

func (bc *Blockchain) MineBlock(message string, txsRootHash []byte, accs map[string]*account.Account) *block.Block {
	newBlock := bc.State[len(bc.State)-1].Block.GenerateBlock()
	newBlock.Data = message
	newBlock.TxsRootHash = txsRootHash
	for i := 0; i < tryLimit; i++ {
		newBlock.Nonce = i
		newBlock.Hash = newBlock.CalculateHash()
		if checkHash(newBlock.Hash, difficulty) {
			break
		}
	}
	return newBlock
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
