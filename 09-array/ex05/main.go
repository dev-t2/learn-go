package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}

	for index, value := range arr1 {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	fmt.Println()

	var arr2 = [5]int{50, 40, 30, 20, 10}

	for index, value := range arr2 {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	fmt.Println()

	arr2 = arr1

	for index, value := range arr2 {
		fmt.Printf("a[%d] = %d\n", index, value)
	}
}