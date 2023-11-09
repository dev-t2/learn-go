package main

import (
	"fmt"
	"reflect"
)

func createPointer() *float64 {
	myFloat := 15.0

	return &myFloat
}

func printPointer(myFloatPointer *float64) {
	fmt.Println(*myFloatPointer)
}

func double(myFloat *float64) {
	*myFloat *= 2
}

func main() {
	myInt := 4

	fmt.Println(myInt)
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt))

	fmt.Println()

	// var myIntPointer *int
	// myIntPointer = &myInt

	// var myIntPointer *int = &myInt

	// var myIntPointer = &myInt

	myIntPointer := &myInt

	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	fmt.Println()

	*myIntPointer = 8 

	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

	fmt.Println()

	myFloatPointer := createPointer()

	printPointer(myFloatPointer)

	double(myFloatPointer)

	fmt.Println(*myFloatPointer)
}