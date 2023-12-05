package main

import (
	"fmt"
	"math"
)

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
)

func main() {
	fmt.Printf("Type: %T\tValue: %v\n", toBe, toBe)
	fmt.Printf("Type: %T\tValue: %v\n", maxInt, maxInt)

	fmt.Println()

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	
	fmt.Println(x, y, z)

	fmt.Println()

	// v := 42
	v := 42.0
	
	fmt.Printf("v is of type %T\n", v)
}