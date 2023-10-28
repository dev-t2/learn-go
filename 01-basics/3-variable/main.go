package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var quantity int
	// var length, width float64
	// var name string

	// quantity = 4
	// length, width = 1.2, 3.75
	// name = "Go"

	// var quantity = 4
	// var length, width = 1.2, 3.75
	// var name = "Go"

	quantity := 4
	length, width := 1.2, 3.75
	name := "Go"

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(name))

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

	height := 2

	fmt.Println(int(width) * height)
	fmt.Println(int(width) > height)
}