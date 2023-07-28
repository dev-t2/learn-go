package main

import (
	"fmt"
)

func Divide1(num1, num2 int) (int, bool) {
	if num2 == 0 {
		return 0, false
	}

	return num1 / num2, true
}

func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false

		// return result, success
		return
	}

	result = a / b
	success = true

	// return result, success
	return
}


func main() {
	var result, success = Divide1(9, 3)

	fmt.Println(result, success)

	result, success = Divide2(9, 0)

	fmt.Println(result, success)
}