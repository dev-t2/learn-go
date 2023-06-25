package main

import "fmt"

func main() {
	var num = 4

	switch num {
	case 1, 2:
		fmt.Println("num == 1 or num == 2")
	case 3, 4, 5:
		fmt.Println("num == 3 or num == 4 or num == 5")
	}
}