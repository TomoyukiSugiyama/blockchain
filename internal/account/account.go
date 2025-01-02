package account

import (
	"strconv"
	"strings"
)

type Account struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func CreateNewAccount(address string, name string, balance int) *Account {
	a := &Account{
		Address: address,
		Name:    name,
		Balance: balance,
	}
	return a
}

func (a *Account) String() string {
	var lines []string
	lines = append(lines, "----- Account -----")
	lines = append(lines, "Address: "+a.Address)
	lines = append(lines, "Name: "+a.Name)
	lines = append(lines, "Balance: "+strconv.Itoa(a.Balance))

	return strings.Join(lines, "\n")
}
