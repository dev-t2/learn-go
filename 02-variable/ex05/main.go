package main

import "fmt"

func main() {
	var num1 float32 = 1234.523
	var num2 float32 = 3456.123
	
	var num3 = num1 * num2 // 4.266663e+06
	var num4 = num3 * 3 // 1.2799989e+07

	fmt.Println(num1, num2, num3, num4)
}