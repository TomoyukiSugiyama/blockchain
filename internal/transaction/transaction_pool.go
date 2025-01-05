package transaction

import (
	"blockchain/utils"
)

type TransactionPool struct {
	utxo []*Transaction
}

func NewTransactionPool() *TransactionPool {
	tp := &TransactionPool{}
	return tp
}

func (tp *TransactionPool) Push(t *Transaction) {
	tp.utxo = append(tp.utxo, t)
}

func (tp *TransactionPool) GetRootHash() []byte {
	var txs [][]byte
	for _, t := range tp.utxo {
		txs = append(txs, t.ToJson())
	}

	root := utils.BuildMerkleTree(txs)

	return root.Hash
}

func (tp *TransactionPool) Pop() []*Transaction {
	t := tp.utxo
	tp.utxo = []*Transaction{}
	return t
}
