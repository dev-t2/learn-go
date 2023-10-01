package main

import "fmt"

func main() {
	nums := [3]float64{71.8, 56.2, 89.5}
	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	length := float64(len(nums))

	fmt.Printf("Average: %0.2f\n", sum / length)
}