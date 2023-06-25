package main

import "fmt"

func main() {
	var num = 3

	switch num {
	case 1:
		fmt.Println("num == 1")
	case 2:
		fmt.Println("num == 2")
	case 3:
		fmt.Println("num == 3")
	case 4:
		fmt.Println("num == 4")
	default: 
		fmt.Println("num > 4")
	}
}