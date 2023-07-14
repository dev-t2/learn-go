package main

import "fmt"

func main() {
	var slice1 = make([]int, 3, 5)
	var slice2 = append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println()

	slice1[1] = 100

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println()

	slice1 = append(slice1, 500)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println()

	var slice3 = []int{1, 2, 3}
	var slice4 = append(slice3, 4, 5)

	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	fmt.Println()

	slice3[1] = 100

	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	fmt.Println()

	slice3 = append(slice3, 500)

	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
}