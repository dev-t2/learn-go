package main

import "fmt"

func three() {
	defer fmt.Println("Deferred in three()")

	panic("This call stack's too deep for me")
}

func two() {
	defer fmt.Println("Deferred in two()")

	three()
}

func one() {
	defer fmt.Println("Deferred in one()")

	two()
}

func main() {
	one()
}