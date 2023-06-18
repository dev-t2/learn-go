package main

import (
	"fmt"
	"math"
	"math/big"
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

	var num4, _ = new(big.Float).SetString("0.1")
	var num5, _ = new(big.Float).SetString("0.2")
	var num6, _ = new(big.Float).SetString("0.3")
	var num7 = new(big.Float).Add(num4, num5)

	fmt.Println(num4, num5, num6, num7)
	fmt.Println(num6.Cmp(num7))
}