package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw")
	}
	a.balance -= amount
	return nil
}
func (a Account) GetBalance() int {
	return a.balance
}

func (a *Account) NewOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint("what ever\nyou like : ", a.GetBalance(), "   ", a.GetOwner())
}
