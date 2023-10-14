package main

import "fmt"

type car struct {
	name 		 string
	topSpeed int
}

func main() {
	var subscriber1 struct {
		name     string
		rate     int
		isActive bool
	}

	subscriber1.name = "Austin"

	fmt.Println("Name:", subscriber1.name)

	var subscriber2 struct {
		name     string
		rate     int
		isActive bool
	}

	subscriber2.name = "Alice"

	fmt.Println("Name:", subscriber2.name)

	fmt.Println()

	var porsche car

	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)
}