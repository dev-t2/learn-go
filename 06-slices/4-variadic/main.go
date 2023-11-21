package main

import (
	"fmt"
	"math"
)

func maximum(nums ...float64) float64 {
	max := math.Inf(-1)

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	floatSlice := []float64{71.8, 56.2, 89.5}

	fmt.Println(maximum(floatSlice...))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
}