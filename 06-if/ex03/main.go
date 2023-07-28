package main

import "fmt"

var count = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", count)

	count++

	return count
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("Count 1 Increase")
	}

	if true || IncreaseAndReturn() < 5 {
		fmt.Println("Count 2 Increase")
	}

	fmt.Println("count:", count)
}