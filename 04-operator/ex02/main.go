package main

import "fmt"

func main() {
	var num1 int8 = 34
	var num2 int16 = 34
	var num3 uint8 = 34
	var num4 uint16 = 34

	fmt.Printf("^%d = %5d,\t %08b\n", num1, ^num1, uint8(^num1))
	fmt.Printf("^%d = %5d,\t %016b\n", num2, ^num2, uint16(^num2))
	fmt.Printf("^%d = %5d,\t %08b\n", num3, ^num3, ^num3)
	fmt.Printf("^%d = %5d,\t %016b\n", num4, ^num4, ^num4)
}