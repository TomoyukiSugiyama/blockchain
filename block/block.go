package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Nonce     int
	Data      string
	PrevHash  string
	Hash      string
}

func (b *Block) Add() {
	// Add block to chain
	fmt.Println("----- Block Added -----")
	fmt.Println("Block Index:", b.Index)
	fmt.Println("Time Stamp:", b.Timestamp)
	fmt.Println("Nonce:", b.Nonce)
	fmt.Println("Data:", b.Data)
	fmt.Println("Previous Hash:", b.PrevHash)
	fmt.Println("Block Hash:", b.Hash)
}

func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) GenerateBlock() *Block {
	newBlock := &Block{
		Index:     b.Index + 1,
		Data:      "",
		Timestamp: time.Now().String(),
		Nonce:     0,
		PrevHash:  b.Hash,
	}
	return newBlock
}
