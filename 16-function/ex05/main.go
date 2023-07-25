package main

import "fmt"

func add(num1, num2 int) int {
	return num1 + num2
}

func mul(num1, num2 int) int {
	return num1 * num2
}

func getOperator(operator string) func(int, int) int {
	if operator == "+" {
		return add
	} else if operator == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int = getOperator("*")
	var result = operator(3, 4)

	fmt.Println(result)
}