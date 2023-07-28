package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str1 = "Hello World"
	var str2 = str1

	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println()

	var stringHeader1 = (*reflect.StringHeader)(unsafe.Pointer(&str1))
	var stringHeader2 = (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}