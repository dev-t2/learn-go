package main

import "fmt"

func main() {	
	if true {
		fmt.Println("true")
	}

	if false {
		fmt.Println("false")
	}

	fmt.Println()
	
	if 1 < 5 {
		fmt.Println("1 < 5")
	}

	if 1 > 5 {
		fmt.Println("1 > 5")
	}

	if 1 <= 5 {
		fmt.Println("1 <= 5")
	}

	if 1 >= 5 {
		fmt.Println("1 >= 5")
	}

	if 1 == 1 {
		fmt.Println("1 == 1")
	}

	if 5 != 5 {
		fmt.Println("5 != 5")
	}

	fmt.Println()

	if !true {
		fmt.Println("!true")
	}

	if !false {
		fmt.Println("!false")
	}

	fmt.Println()

	if true && 1 == 1 {
		fmt.Println("true && 1 == 1")
	}

	if false && 1 == 1 {
		fmt.Println("false && 1 == 1")
	}

	if true || 1 != 1 {
		fmt.Println("true || 1 != 1")
	}

	if false || 1 != 1 {
		fmt.Println("false || 1 != 1")
	}
}