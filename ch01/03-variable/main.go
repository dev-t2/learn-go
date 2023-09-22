package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var quantity int
	// var length, width float64
	// var customerName string

	// quantity = 4
	// length, width = 1.2, 3.75
	// customerName = "Austin"

	// var quantity = 4
	// var length, width = 1.2, 3.75
	// var customerName = "Austin"

	quantity := 4
	length, width := 1.2, 3.75
	customerName := "Austin"

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))

	fmt.Println()

	var myInt int
	var myFloat float64
	var myString string
	var myBool bool

	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myString)
	fmt.Println(myBool)

	fmt.Println()

	myInt = 2

	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))

	fmt.Println()

	height := 2

	fmt.Println(width * float64(height))
	fmt.Println(width > float64(height))
	fmt.Println(int(width))
}