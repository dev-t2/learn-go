package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	fmt.Println(target)
}