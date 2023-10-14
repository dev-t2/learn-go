package main

import "fmt"

type subscriber struct {
	name     string
	rate     int
	isActive bool
}

func main() {
	var subscriber1 subscriber

	subscriber1.name = "Austin"

	fmt.Println("Name:", subscriber1.name)

	var subscriber2 subscriber

	subscriber2.name = "Alice"

	fmt.Println("Name:", subscriber2.name)
}