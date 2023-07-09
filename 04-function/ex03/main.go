package main

import "fmt"

func Divide1(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false

		return
	}

	result = a / b
	success = true

	return
}

func main() {
	var result, success = Divide1(9, 3)

	fmt.Println(result, success)

	result, success = Divide2(9, 0)

	fmt.Println(result, success)
}