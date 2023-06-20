package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	var num, success = Divide(9, 3)

	fmt.Println(num, success)

	num, success = Divide(9, 0)

	fmt.Println(num, success)
}