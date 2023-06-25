package main

import "fmt"

func getNumber() int {
	return 22
}

func main() {
	switch num := getNumber(); num {
	case 10:
		fmt.Println("num == 10")
	case 33:
		fmt.Println("num == 33")
	default:
		fmt.Println("num:", num)
	}

	switch num := getNumber(); true {
	// switch num := getNumber(); {
	case num < 10:
		fmt.Println("num < 10")
	case num < 20:
		fmt.Println("num < 20")
	case num < 30:
		fmt.Println("num < 30")
	default:
		fmt.Println("num:", num)
	}
}