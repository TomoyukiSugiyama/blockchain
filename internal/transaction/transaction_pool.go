package transaction

type TransactionPool struct {
	Transactions []*Transaction
}

func NewTransactionPool() *TransactionPool {
	tp := &TransactionPool{}
	return tp
}

func (tp *TransactionPool) Push(t *Transaction) {
	tp.Transactions = append(tp.Transactions, t)
}

func (tp *TransactionPool) Pop() (x *Transaction) {
	if len(tp.Transactions) == 0 {
		return nil
	}
	x, tp.Transactions = tp.Transactions[0], tp.Transactions[1:]
	return x
}
