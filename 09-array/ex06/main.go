package main

import "fmt"

func main() {
	var arrays = [2][5]int{ {1, 2, 3, 4, 5}, {6, 7, 8, 9, 10} }

	for _, array := range arrays {
		for _, value := range array {
			fmt.Print(value, " ")
		}

		fmt.Println()
	}
}