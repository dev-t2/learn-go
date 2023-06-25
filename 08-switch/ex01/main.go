package main

import "fmt"

func main() {
	var num = 3

	if num == 1 {
		fmt.Println("If: num == 1")
	} else if num == 2 {
		fmt.Println("If: num == 2")
	} else if num == 3 {
		fmt.Println("If: num == 3")
	} else if num == 4 {
		fmt.Println("If: num == 3")
	} else {
		fmt.Println("If: num > 4")
	}

	switch num {
	case 1:
		fmt.Println("Switch: num == 1")
	case 2:
		fmt.Println("Switch: num == 2")
	case 3:
		fmt.Println("Switch: num == 3")
	case 4:
		fmt.Println("Switch: num == 4")
	default: 
		fmt.Println("Switch: num > 4")
	}
}