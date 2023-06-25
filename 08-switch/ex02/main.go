package main

import "fmt"

func main() {
	var num = 19

	switch true {
	case num < 10 || num > 30:
		fmt.Println("num < 10 || num >= 30")
	case num >= 10 && num < 20:
		fmt.Println("num >= 10 && num < 20")
	case num >= 15 && num < 25:
		fmt.Println("num >= 15 && num < 25")
	default:
		fmt.Println("num >= 25 && num < 30")
	}
}