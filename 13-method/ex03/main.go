package main

import "fmt"

type account struct {
	balance int
}

func (account *account) withdrawPointer(amount int) {
	account.balance -= amount
}

func (account account) withdrawValue(amount int) {
	account.balance -= amount
}

func (account account) withdrawReturnValue(amount int) account {
	account.balance -= amount

	return account
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