package main

import "fmt"

func main() {
	var num1 = 0.1
	var num2 = 0.2
	var num3 = 0.3

	fmt.Println(num1 + num2)
	fmt.Printf("%f + %f == %f: %v", num1, num2, num3, num1 + num2 == num3)
}