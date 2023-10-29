package main

import (
	"fmt"
	"math/rand"
)

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

func awardPrize() {
	doorNumber := rand.Intn(3) + 1

	if doorNumber == 1 {
		fmt.Println("You Win a house")
	} else if doorNumber == 2 {
		fmt.Println("You Win a car")
	} else if doorNumber == 3 {
		fmt.Println("You Win a notebook")
	} else {
		panic("Invalid door number")
	}
}

func main() {
	// one()

	awardPrize()
}