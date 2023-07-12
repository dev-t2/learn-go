package main

import "fmt"

func updateArray(array [5]int) {
	array[2] = 20
}

func updateSlice(slice []int) {
	slice[2] = 20
}

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}

	updateArray(array)
	updateSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}