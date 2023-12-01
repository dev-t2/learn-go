package main

import (
	"fmt"
	"reflect"
)

func createPointer() *int {
	myInt := 15

	return &myInt
}

func printPointer(myInt *int) {
	fmt.Println(*myInt)
}

func double(myInt *int) {
	*myInt *= 2
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

	myIntPointer = createPointer()

	printPointer(myIntPointer)

	double(myIntPointer)

	fmt.Println(*myIntPointer)
}