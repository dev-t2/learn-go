package main

import "fmt"

type operatorType func(num1, num2 int) int

func getOperator(operator string) operatorType {
	if operator == "+" {
		return func(num1, num2 int) int {
			return num1 + num2
		}
	}

	if operator == "*" {
		return func(num1, num2 int) int {
			return num1 * num2
		}
	}

	return nil
}

func main() {
	var operator = getOperator("*")
	var result = operator(3, 4)

	fmt.Println(result)
}