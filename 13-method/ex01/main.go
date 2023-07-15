package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(account *account, amount int) {
	account.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	var account = &account{100}

	withdrawFunc(account, 30)

	fmt.Printf("%d\n", account.balance)

	account.withdrawMethod(30)

	fmt.Printf("%d\n", account.balance)
}