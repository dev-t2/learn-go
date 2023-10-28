package main

import (
	"fmt"
	"learn-go/09-structs/user"
)

func main() {
	subscriber1 := user.DefaultSubscriber("Austin")
	
	user.PrintInfo(subscriber1)

	fmt.Println()

	subscriber2 := user.DefaultSubscriber("Alice")

	user.ApplyDiscount(subscriber2)

	user.PrintInfo(subscriber2)

	fmt.Println()

	subscriber := user.Subscriber{Name: "Chloe"}
	address := user.Address{Street: "Street", City: "City", State: "State"}
	
	subscriber.HomeAddress = address
	subscriber.HomeAddress.PostalCode = "12345"

	fmt.Printf("%#v\n", subscriber)

	fmt.Println()

	employee := user.Employee{Name: "Sally", Salary: 5000}

	employee.State = "Street"
	employee.City = "City"
	employee.State = "State"
	employee.PostalCode = "12345"

	fmt.Printf("%#v\n", employee)
}