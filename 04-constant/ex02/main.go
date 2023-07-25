package main

import "fmt"

func main() {
	const PI1 = 3.141592653589793238

	var PI2 = 3.141592653589793238

	PI2 = 3

	fmt.Printf("원주율: %f\n", PI1)
	fmt.Printf("원주율: %f\n", PI2)
}