package main

import (
	"fmt"
	"learn-go/ch08/packages/magazine"
)

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber

	s.Name = name
	s.Rate = 9900
	s.IsActive = true

	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4900
}

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly Rate:", s.Rate)
	fmt.Println("Active:", s.IsActive)
}

func main() {
	subscriber1 := defaultSubscriber("Austin")
	
	applyDiscount(subscriber1)
	
	printInfo(subscriber1)

	fmt.Println()

	subscriber2 := defaultSubscriber("Alice")

	printInfo(subscriber2)
}