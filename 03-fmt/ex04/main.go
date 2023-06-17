package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	var num, err = fmt.Scan(&num1, &num2)

	if err == nil {
		fmt.Println(num, num1, num2)
	} else {
		fmt.Println(num, err)
	}
}