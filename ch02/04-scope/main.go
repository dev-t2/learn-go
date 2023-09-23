package main

import "fmt"

var packageVar = "package"

func main() {
	functionVar := "function"

	if true {
		conditionalVar := "conditional"

		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionalVar)
	}

	fmt.Println()

	fmt.Println(packageVar)
	fmt.Println(functionVar)
	// fmt.Println(conditionalVar)
}