package ledger

import "github.com/app/utilities"

type User struct {
	ID   string
	Name string
}

func NewUser(name string) *User {
	u := &User{
		ID:   utilities.GenerateUUID(),
		Name: name,
	}
	AddUser(u)
	return u
}
