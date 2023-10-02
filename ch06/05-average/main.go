package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(nums ...float64) float64 {
	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

func main() {
	arguments := os.Args[1:]
	
	var nums []float64

	for _, argument := range arguments {
		num, err := strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(err)
		}
		
		nums = append(nums, num)
	}

	fmt.Printf("Average: %0.2f\n", average(nums...))
}