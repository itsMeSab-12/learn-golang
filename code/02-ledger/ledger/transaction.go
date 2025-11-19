package ledger

import (
	"time"

	"github.com/app/utilities"
)

type Transaction struct {
	ID     string
	From   *User
	To     *User
	Amount float32
	At     string
}

func NewTransaction(a *User, b *User, amt float32) *Transaction {

	tx := &Transaction{
		ID:     utilities.GenerateUUID(),
		From:   a,
		To:     b,
		Amount: amt,
		At:     time.Now().Format(time.RFC3339),
	}

	AddTransaction(tx)

	return tx
}
