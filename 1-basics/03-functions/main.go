package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println()

	a, b := swap("hello", "world")

	fmt.Println(a, b)

	fmt.Println()

	fmt.Println(split(17))
}