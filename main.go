package main

import (
	accounts "cencivic/learngo/accounts/acccounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("kyunbin")
	err := account.Deposit(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.GetBalance())
	fmt.Println(account)

}
