package main

import "fmt"

type subscriber struct {
	name     string
	rate     int
	isActive bool
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber

	s.name = name
	s.rate = 9900
	s.isActive = true

	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4900
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly Rate:", s.rate)
	fmt.Println("Active:", s.isActive)
}

func main() {
	subscriber1 := defaultSubscriber("Austin")
	
	applyDiscount(subscriber1)
	
	printInfo(subscriber1)

	fmt.Println()

	subscriber2 := defaultSubscriber("Alice")

	printInfo(subscriber2)
}