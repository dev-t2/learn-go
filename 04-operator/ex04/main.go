package main

import "fmt"

func main() {
	var num1 int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", num1, num1, num1 < num1 + 1)
	fmt.Printf("num1     = %4d,  %08b\n", num1, num1)
	fmt.Printf("num1 + 1 = %4d, %08b\n", num1 + 1, num1 + 1)

	var num2 int8 = -128
	
	fmt.Printf("%d > %d - 1: %v\n", num2, num2, num2 > num2 - 1)
	fmt.Printf("num2     = %4d, %08b\n", num2, num2)
	fmt.Printf("num2 - 1 = %4d,  %08b\n", num2 - 1, num2 - 1)
}