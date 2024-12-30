package block

import (
	"strconv"
	"strings"
)

type Transaction struct {
	Id     int
	From   string
	To     string
	Amount int
}

func CreateNewTransaction(id int, from, to string, amount int) *Transaction {
	t := &Transaction{
		Id:     id,
		From:   from,
		To:     to,
		Amount: amount,
	}
	return t
}

func (t *Transaction) String() string {
	var lines []string
	lines = append(lines, "----- Transaction -----")
	lines = append(lines, "Transaction Id: "+strconv.Itoa(t.Id))
	lines = append(lines, "From: "+t.From)
	lines = append(lines, "To: "+t.To)
	lines = append(lines, "Amount: "+strconv.Itoa(t.Amount))

	return strings.Join(lines, "\n")
}
