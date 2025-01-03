package transaction

type TransactionPool struct {
	Transactions []*Transaction
}

func NewTransactionPool() *TransactionPool {
	tp := &TransactionPool{}
	return tp
}
