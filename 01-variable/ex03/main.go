package main

import "fmt"

func main() {
	var num1 = 1
	var num2 = 2.5

	var num3 int = int(num2)
	var num4 float64 = float64(num1 * num3)

	var num5 int64 = 5
	var num6 int64 = int64(num4) * num5
	
	var num7 int = int(num2 * 7)
	var num8 int = int(num2) * 8

	fmt.Println(num6, num7, num8)

	var num9 int16 = 3456
	var num10 int8 = int8(num9)

	fmt.Println(num9, num10)
}