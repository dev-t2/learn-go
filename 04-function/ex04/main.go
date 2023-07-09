package main

import "fmt"

func printNumber(num int) {
	if num == 0 {
		return
	}

	fmt.Println(num)

	printNumber(num - 1)

	fmt.Println("After", num)
}

func main() {
	printNumber(3)
}