package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6

	fmt.Println(amount)
	fmt.Println(&amount)

	fmt.Println()

	var myInt int

	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt))

	fmt.Println()

	// var myIntPointer *int
	// myIntPointer = &myInt

	myIntPointer := &myInt

	fmt.Println(myIntPointer)
}