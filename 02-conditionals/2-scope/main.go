package main

import "fmt"

var variable1 = "Package Variable"

func main() {
	variable2 := "Function Variable"

	if true {
		variable3 := "Conditional Variable"

		fmt.Println(variable1)
		fmt.Println(variable2)
		fmt.Println(variable3)
	}

	fmt.Println()

	fmt.Println(variable1)
	fmt.Println(variable2)
	// fmt.Println(variable3)

	fmt.Println()

	// count := 5

	// if count > 1 {
	// 	fmt.Println(count)
	// }

	// fmt.Println(count)

	if count := 5; count > 1 {
		fmt.Println(count)
	}

	// fmt.Println(count)
}