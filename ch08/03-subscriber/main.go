package main

import "fmt"

func main() {
	var subscriber struct {
		name     string
		rate     float64
		isActive bool
	}

	subscriber.name = "Austin"
	subscriber.rate = 9900
	subscriber.isActive = true

	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly Rate:", subscriber.rate)
	fmt.Println("Active:", subscriber.isActive)
}