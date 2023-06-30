package main

import "fmt"

const Len = 3

func main() {
	var len = 5

	// var arr1 = [len]int{1, 2, 3, 4, 5}

	fmt.Println(len)

	var arr2 = [Len]int{1, 2, 3}

	fmt.Println(arr2)

	var arr3 [5]int

	fmt.Println(arr3)
}