package main

import "fmt"

func sum(numbers ...int) int {
	var result = 0

	fmt.Printf("Numbers Type: %T\n", numbers)

	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(sum(10, 20))

	fmt.Println(sum())
}