package state

import (
	"blockchain/internal/account"
	"blockchain/internal/block"
	"strings"

	"encoding/json"
)

type State struct {
	Accounts map[string]*account.Account `json:"accounts"`
	Block    *block.Block                `json:"block"`
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

func (s *State) ToJson() []byte {
	jsonData, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func (s *State) FromJson(data []byte) {
	err := json.Unmarshal(data, s)
	if err != nil {
		panic(err)
	}
}
