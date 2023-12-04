package main

import "fmt"

var bool1, bool2 bool

var num1, num2 int = 1, 2

func main() {
	var num3 int
	
	fmt.Println(num3, bool1, bool2)

	fmt.Println()

	var bool3, str1 = true, "string"

	fmt.Println(num1, num2, bool3, str1)

	fmt.Println()

	var num4, num5 int = 1, 2
	
	num6 := 3
	bool4, str2 := true, "string"

	fmt.Println(num4, num5, num6, bool4, str2)
}