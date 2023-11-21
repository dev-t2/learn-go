package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(nums ...float64) (result float64) {
	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

func main() {
	args := os.Args[1:]
	
	var nums []float64

	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			log.Fatal(err)
		}
		
		nums = append(nums, num)
	}

	result := average(nums...)

	fmt.Printf("Average: %0.2f\n", result)
}