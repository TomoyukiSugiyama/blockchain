package account

import (
	"strconv"
	"strings"
)

type Account struct {
	Id      string
	Name    string
	Balance int
}

func CreateNewAccount(id string, name string, balance int) *Account {
	a := &Account{
		Id:      id,
		Name:    name,
		Balance: balance,
	}
	return a
}

func (a *Account) String() string {
	var lines []string
	lines = append(lines, "----- Account -----")
	lines = append(lines, "Account Id: "+a.Id)
	lines = append(lines, "Name: "+a.Name)
	lines = append(lines, "Balance: "+strconv.Itoa(a.Balance))

	return strings.Join(lines, "\n")
}
