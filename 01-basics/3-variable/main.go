package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myFloat float64
	var myString string
	var myBool bool

	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myString)
	fmt.Println(myBool)

	fmt.Println()

	// var num1 int
	// num1 = 1

	// var num1 int = 1

	// var num1 = 1

	num1 := 1
	num2 := 5.0

	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num2))

	fmt.Println()

	fmt.Println(num1 + int(num2))
	fmt.Println(num1 < int(num2))

	fmt.Println()

	fmt.Println(int(num2) == 5)
	fmt.Println(int(num2) == 5.0)
}