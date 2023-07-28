package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str = "Hello World"
	var bytes = []byte(str)

	bytes[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", bytes)

	fmt.Println()

	bytes[2] = 'l'

	var strHeader = (*reflect.StringHeader)(unsafe.Pointer(&str))
	var bytesHeader = (*reflect.StringHeader)(unsafe.Pointer(&bytes))

	fmt.Printf("str:\t%x\n", strHeader.Data)
	fmt.Printf("bytes:\t%x\n", bytesHeader.Data)
}