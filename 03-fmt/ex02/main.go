package main

import "fmt"

func main() {
	var num1 = 123
	var num2 = 456
	var num3 = 123456789

	fmt.Printf("%5d, %5d, %5d\n", num1, num2, num3)
	fmt.Printf("%05d, %05d, %05d\n", num1, num2, num3)
	fmt.Printf("%-5d, %-5d, %-5d\n", num1, num2, num3)
}