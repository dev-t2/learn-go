package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part

	p.description = description
	p.count = 100

	return p
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber) {
	// (*s).rate = 4.99
	s.rate = 4.99
}

func main() {
	var bolts part

	bolts.description = "Bolts"
	bolts.count = 25

	showInfo(bolts)

	fmt.Println()

	p := minimumOrder("Bolts")

	fmt.Println(p.description, p.count)

	fmt.Println()

	var s subscriber

	fmt.Println(s.rate)

	applyDiscount(&s)

	fmt.Println(s.rate)
}