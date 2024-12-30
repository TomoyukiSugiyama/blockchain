package blockchain

import (
	"blockchain/account"
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

	// Test adding accounts
	acc1 := account.CreateNewAccount("0000", "Alice", 1000)
	acc2 := account.CreateNewAccount("0001", "Bob", 1000)
	accs := map[string]*account.Account{acc1.Id: acc1, acc2.Id: acc2}
	fmt.Println(acc1.String())
	fmt.Println(acc2.String())
	// Test adding transactions
	tr1 := block.CreateNewTransaction(0, acc1.Id, acc2.Id, 100)
	tr2 := block.CreateNewTransaction(1, acc1.Id, acc2.Id, 10)
	trs := []block.Transaction{*tr1, *tr2}
	// Test adding a block
	bc.MineBlock("First Block", trs, accs)
	fmt.Println(acc1.String())
	fmt.Println(acc2.String())

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
	fmt.Println(bc.Blocks[genesisBlock.Hash].String())
}

func (bc *Blockchain) addBlock(b *block.Block, h string) {
	if !checkHash(h, difficulty) {
		// Reject block
		fmt.Println("Block Rejected with hash: ", h)
		return
	}

	if !isExistPreviousBlock(bc.Blocks, b.PrevHash) {
		// Reject block
		fmt.Println("Block Rejected")
		return
	}

	bc.Blocks[h] = b
	fmt.Println(bc.Blocks[h].String())
}

func (bc *Blockchain) MineBlock(message string, trs []block.Transaction, accs map[string]*account.Account) {
	newBlock := bc.currentBlock.GenerateBlock()
	newBlock.Data = message
	newBlock.Transactions = trs
	for i := 0; i < tryLimit; i++ {
		newBlock.Nonce = i
		newBlock.Hash = newBlock.CalculateHash()
		if checkHash(newBlock.Hash, difficulty) {
			break
		}
	}
	bc.addBlock(newBlock, newBlock.Hash)
	for _, t := range trs {
		t.Run(accs)
	}
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
