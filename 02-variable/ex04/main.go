package main

import "fmt"

var num1 = 10

func main() {
	var num2 = 20
	var num3 = 30

	{
		var num3 = 40

		fmt.Println(num1, num2, num3)
	}

	fmt.Println(num3)
}