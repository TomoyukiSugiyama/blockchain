package state

import (
	"blockchain/internal/account"
	"blockchain/internal/block"
	"strings"
)

type State struct {
	Accounts map[string]*account.Account
	Block    *block.Block
}

func CreateNewState(a map[string]*account.Account, b *block.Block) *State {
	return &State{Accounts: a, Block: b}
}

func (s *State) String() string {
	var lines []string
	lines = append(lines, "----- State -----")
	lines = append(lines, s.Block.String())
	for _, acc := range s.Accounts {
		lines = append(lines, acc.String())
	}

	return strings.Join(lines, "\n")
}
