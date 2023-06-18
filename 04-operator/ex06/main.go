package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var num1 = 0.1
	var num2 = 0.2
	var num3 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", num1, num2, num1 + num2)
	fmt.Printf("%0.18f == %0.18f: %v\n", num1 + num2, num3, equal(num1 + num2, num3))

	num1 = 0.0000000000004
	num2 = 0.0000000000002
	num3 = 0.0000000000007

	fmt.Printf("%g == %g: %v\n", num1 + num2, num3, equal(num1 + num2, num3))
}