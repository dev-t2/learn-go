package main

import "fmt"

func find45(num int) (int, bool) {
	for i := 1; i <= 9; i++ {
		if num*i == 45 {
			return i, true
		}
	}

	return 0, false
}

func main() {
	var num1, num2 = 1, 0

	for ; num1 <= 9; num1++ {
		var isExit = false

		if num2, isExit = find45(num1); isExit {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", num1, num2, num1 * num2)
}