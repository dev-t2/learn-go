package main

import (
	"fmt"
	"learn-go/12-interfaces/mypackage"
)

func main() {
	var value mypackage.MyInterface

	value = mypackage.MyInt(5)

	value.MethodWithoutParameter()
	value.MethodWithParameter(123.4)
	fmt.Println(value.MethodWithReturn())
}