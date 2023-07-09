package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	var num1 = Add(1, 2)

	fmt.Println(num1)
}