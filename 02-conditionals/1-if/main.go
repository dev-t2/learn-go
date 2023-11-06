package main

import "fmt"

func main() {	
	if true {
		fmt.Println("true")
	}

	if false {
		fmt.Println("false")
	}
	
	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	if 1 > 2 {
		fmt.Println("1 > 2")
	}

	if 1 <= 2 {
		fmt.Println("1 <= 2")
	}

	if 1 >= 2 {
		fmt.Println("1 >= 2")
	}

	if 1 == 1 {
		fmt.Println("1 == 1")
	}

	if 2 != 2 {
		fmt.Println("2 != 2")
	}

	if !true {
		fmt.Println("!true")
	}

	if !false {
		fmt.Println("!false")
	}

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