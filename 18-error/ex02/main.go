package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Error: %g", f)
	}

	return math.Sqrt(f), nil
}

func main() {
	var sqrt, error = Sqrt(-2)

	if error != nil {
		fmt.Printf("%v\n", error)

		return
	}

	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}