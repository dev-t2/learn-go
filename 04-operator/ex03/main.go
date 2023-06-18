package main

import "fmt"

func main() {
	var num1 int8 = 4

	fmt.Printf("num1: %08b, num1 << 2: %08b, num1 << 2: %d\n", num1, num1 << 2, num1 << 2)

	var num2 int8 = 64

	fmt.Printf("num2: %08b, num2 << 2: %08b, num2 << 2: %d\n", num2, num2 << 2, num2 << 2)

	var num3 int8 = 16

	fmt.Printf("num3: %08b, num3 >> 2: %08b, num3 >> 2: %d\n", num3, num3 >> 2, num3 >> 2)

	var num4 int8 = -128

	fmt.Printf("num4: %08b, num4 >> 2: %08b, num4 >> 2: %d\n", uint8(num4), uint8(num4 >> 2), num4 >> 2)

	var num5 int8 = -1

	fmt.Printf("num5: %08b, num5 >> 2: %08b, num5 >> 2: %d\n", uint8(num5), uint8(num5 >> 2), num5 >> 2)

	var num6 uint8 = 128

	fmt.Printf("num6: %08b, num6 >> 2: %08b, num6 >> 2: %d\n", uint8(num6), uint8(num6 >> 2), num6 >> 2)
}