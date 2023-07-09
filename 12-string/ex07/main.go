package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func ToUpper(str string) string {
	var builder strings.Builder

	for _, value := range str {
		if value >= 'a' && value <= 'z' {
			builder.WriteRune('A' + (value - 'a'))
		} else {
			builder.WriteRune(value)
		}
	}

	return builder.String()
}

func main() {
	var str = "Hello"
	var strHeader = (*reflect.StringHeader)(unsafe.Pointer(&str))
	var addr1 = strHeader.Data

	str += " Go"

	var addr2 = strHeader.Data

	str += " World"

	var addr3 = strHeader.Data

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)

	fmt.Println()

	var addr4 = strHeader.Data

	fmt.Println(ToUpper(str))
	fmt.Printf("addr4:\t%x\n", addr4)
}