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
	// length, width = 1.2, 3.4
	// customerName = "Austin"

	var quantity = 4
	var length, width = 1.2, 3.4
	var customerName = "Austin"

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))
}