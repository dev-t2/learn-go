package main

import "fmt"

func main() {
	var num1 int32 = 1
	var num2 int32 = 2

	fmt.Println("num1 + num2 =", num1 + num2)
	fmt.Println("num1 - num2 =", num1 - num2)
	fmt.Println("num1 * num2 =", num1 * num2)
	fmt.Println("num1 / num2 =", num1 / num2)
	fmt.Println("num1 % num2 =", num1 % num2)

	var num3 float32 = 3.14
	var num4 float32 = 4

	fmt.Println("num3 + num4 =", num3 + num4)
	fmt.Println("num3 - num4 =", num3 - num4)
	fmt.Println("num3 * num4 =", num3 * num4)
	fmt.Println("num3 / num4 =", num3 / num4)
}