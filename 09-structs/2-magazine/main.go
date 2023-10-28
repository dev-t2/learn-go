package main

import (
	"fmt"
	"learn-go/09-structs/packages/magazine"
)

func main() {
	subscriber1 := magazine.DefaultSubscriber("Austin")
	
	magazine.PrintInfo(subscriber1)

	fmt.Println()

	subscriber2 := magazine.DefaultSubscriber("Alice")

	magazine.ApplyDiscount(subscriber2)

	magazine.PrintInfo(subscriber2)

	fmt.Println()

	subscriber := magazine.Subscriber{Name: "Chloe"}
	address := magazine.Address{Street: "Street", City: "City", State: "State"}
	
	subscriber.HomeAddress = address
	subscriber.HomeAddress.PostalCode = "12345"

	fmt.Printf("%#v\n", subscriber)

	fmt.Println()

	employee := magazine.Employee{Name: "Sally", Salary: 5000}

	employee.State = "Street"
	employee.City = "City"
	employee.State = "State"
	employee.PostalCode = "12345"

	fmt.Printf("%#v\n", employee)
}