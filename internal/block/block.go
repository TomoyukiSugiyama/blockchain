package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index       int    `json:"index"`
	Timestamp   string `json:"timestamp"`
	Nonce       int    `json:"nonce"`
	Data        string `json:"data"`
	PrevHash    string `json:"prevHash"`
	Hash        string `json:"hash"`
	TxsRootHash []byte `json:"txsRootHash"`
}

func (b *Block) String() string {
	var lines []string
	lines = append(lines, "----- Block -----")
	lines = append(lines, "Block Index: "+strconv.Itoa(b.Index))
	lines = append(lines, "Time Stamp: "+b.Timestamp)
	lines = append(lines, "Nonce: "+strconv.Itoa(b.Nonce))
	lines = append(lines, "Data: "+b.Data)
	lines = append(lines, "Previous Hash: "+b.PrevHash)
	lines = append(lines, "Block Hash: "+b.Hash)
	lines = append(lines, "TxsRootHash: "+hex.EncodeToString(b.TxsRootHash))

	return strings.Join(lines, "\n")
}

func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce) + hex.EncodeToString(b.TxsRootHash)
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

func (b *Block) ToJson() []byte {
	jsonData, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func (b *Block) FromJson(data []byte) {
	err := json.Unmarshal(data, b)
	if err != nil {
		panic(err)
	}
}
