package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	var n, err = fmt.Scanln(&num1, &num2)

	if err == nil {
		fmt.Println(n, num1, num2)
	} else {
		fmt.Println(n, err)
	}
}