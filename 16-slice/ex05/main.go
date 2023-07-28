package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = make([]int, len(slice1))

	for index, value := range slice1 {
		slice2[index] = value
	}

	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println()

	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = append([]int{}, slice3...)

	slice3[1] = 100

	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Println()

	var slice5 = []int{1, 2, 3, 4, 5}
	var slice6 = make([]int, 3, 10)
	var slice7 = make([]int, 10)

	var count1 = copy(slice6, slice5)
	var count2 = copy(slice7, slice5)

	fmt.Println(slice5)
	fmt.Println(slice6, count1)
	fmt.Println(slice7, count2)
}