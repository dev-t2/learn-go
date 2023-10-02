package main

import (
	"fmt"
	"learn-go/packages/datafile"
	"log"
)

func main() {
	nums, err := datafile.GetFloats("ch06/04-average/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	length := float64(len(nums))

	fmt.Printf("Average: %0.2f\n", sum / length)
}