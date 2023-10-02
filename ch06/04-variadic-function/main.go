package main

import (
	"fmt"
	"math"
)

func severalInts(nums ...int) {
	fmt.Println(nums)
}

func maximum(nums ...float64) float64 {
	max := math.Inf(-1)

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func isRange(min, max float64, nums ...float64) []float64 {
	var results []float64

	for _, num := range nums {
		if num >= min && num <= max {
			results = append(results, num)
		}
	}

	return results
}

func main() {
	severalInts(1, 2, 3)

	intSlice := []int{4, 5, 6}

	severalInts(intSlice...)

	fmt.Println()

	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))

	fmt.Println()

	fmt.Println(isRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(isRange(-10, 10, 4.1, 12, -12, -5.2))
}