package main

import (
	"fmt"
	"learn-go/ch08/packages/magazine"
)

func main() {
	subscriber1 := magazine.DefaultSubscriber("Austin")
	
	magazine.ApplyDiscount(subscriber1)
	
	magazine.PrintInfo(subscriber1)

	fmt.Println()

	subscriber2 := magazine.DefaultSubscriber("Alice")

	magazine.PrintInfo(subscriber2)

	fmt.Println()

	subscriber := magazine.Subscriber{Name: "Cloyee"}
	address := magazine.Address{Street: "Street", City: "City", State: "State"}
	
	subscriber.HomeAddress = address

	fmt.Printf("%#v\n", subscriber)

	fmt.Println()

	employee := magazine.Employee{Name: "Sally", Salary: 5000}

	employee.HomeAddress.State = "Street"
	employee.HomeAddress.City = "City"
	employee.HomeAddress.State = "State"

	fmt.Printf("%#v\n", employee)
}