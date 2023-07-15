package main

import "fmt"

type account struct {
	balance int
}

func (a *account) withdrawPointer(amount int) {
	a.balance -= amount
}

func (a account) withdrawValue(amount int) {
	a.balance -= amount
}

func (a account) withdrawReturnValue(amount int) account {
	a.balance -= amount

	return a
}

func main() {
	var account1 *account = &account{100}

	account1.withdrawPointer(30)

	fmt.Println(account1.balance)

	account1.withdrawValue(20)

	fmt.Println(account1.balance)

	var account2 account = account1.withdrawReturnValue(20)

	fmt.Println(account2.balance)

	account2.withdrawPointer(30)

	fmt.Println(account2.balance)
}