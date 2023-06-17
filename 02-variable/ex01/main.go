package main

import "fmt"

func main() {
	var num1 int = 1
	var str1 string = "a"

	num1 = 2
	str1 = "b"

	fmt.Println(num1, str1)

	var num2 int = 3
	var num3 int = 4
	var sum int = num2 + num3

	fmt.Println(sum)
}