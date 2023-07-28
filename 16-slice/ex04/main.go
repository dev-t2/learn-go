package main

import "fmt"

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	var slice = array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	fmt.Println()

	array[1] = 100

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	fmt.Println()

	slice = append(slice, 500)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}