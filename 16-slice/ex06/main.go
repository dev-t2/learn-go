package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6}
	var index = 2

	for i := index + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}

	slice = slice[:len(slice)-1]

	fmt.Println(slice)

	fmt.Println()

	slice = []int{1, 2, 3, 4, 5, 6}

	slice = append(slice[:index], slice[index+1:]... )

	fmt.Println(slice)
}