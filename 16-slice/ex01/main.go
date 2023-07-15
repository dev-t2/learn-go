package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println(slice)
	}

	slice = append(slice, 0)

	fmt.Println(slice)

	for i := 0; i < 10; i++ {
		slice = append(slice, i + 1)
	}

	fmt.Println(slice)

	slice = append(slice, 11, 12, 13, 14, 15)

	fmt.Println(slice)
}