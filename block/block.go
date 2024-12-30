package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Block struct {
	Index     int
	Timestamp string
	Nonce     int
	Data      string
	PrevHash  string
	Hash      string
}

func Add(b *Block) {
	// Add block to chain
	fmt.Println("----- Block Added -----")
	fmt.Println("Block Index:", b.Index)
	fmt.Println("Time Stamp:", b.Timestamp)
	fmt.Println("Nonce:", b.Nonce)
	fmt.Println("Transaction Data:", b.Data)
	fmt.Println("Previous Hash:", b.PrevHash)
}

func CalculateHash(b *Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
