package transaction

import (
	"blockchain/internal/account"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type Transaction struct {
	Id     int    `json:"id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
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

func (t *Transaction) Run(accounts map[string]*account.Account) {
	from := accounts[t.From]
	to := accounts[t.To]
	if from == nil || to == nil {
		log.Fatalln("Invalid transaction")
		return
	}
	if from.Balance < t.Amount {
		log.Fatalln("Insufficient balance")
		return
	}
	from.Balance -= t.Amount
	to.Balance += t.Amount
}

func (t *Transaction) Bytes() []byte {
	jsonData, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func (t *Transaction) FromJson(data []byte) {
	err := json.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}
}
