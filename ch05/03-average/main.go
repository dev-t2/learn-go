package main

import (
	"fmt"
	"learn-go/ch05/03-average/datafile"
	"log"
)

func main() {
	nums, err := datafile.GetFloats("ch05/03-average/data.txt")

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