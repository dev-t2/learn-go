package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)

	var index = 2

	for i := len(slice) - 2; i >= index; i-- {
		slice[i+1] = slice[i]
	}

	slice[index] = 100

	fmt.Println(slice)

	fmt.Println()

	slice = []int{1, 2, 3, 4, 5, 6}

	slice = append(slice[:index], append([]int{100}, slice[index:]...)...)

	fmt.Println(slice)

	fmt.Println()

	slice = []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)

	copy(slice[index+1:], slice[index:])

	slice[index] = 100

	fmt.Println(slice)
}