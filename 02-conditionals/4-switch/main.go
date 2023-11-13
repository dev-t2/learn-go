package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch rand.Intn(3) + 1 {
	case 1:
		fmt.Println("Number 1")
	case 2:
		fmt.Println("Number 2")
	case 3:
		fmt.Println("Number 3")
	default:
		panic("Invalid Number")		
	}
}