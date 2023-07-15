package main

import (
	"fmt"
	"learn-go/14-package/ex03/custom"
)

func main() {
	fmt.Println("Public Constant:", custom.PublicConstant)

	fmt.Println("Public Variable:", custom.PublicVariable)

	custom.PublicFunction()

	var publicType custom.PublicType = 1

	fmt.Println("Public Type:", publicType)

	var publicStruct = custom.PublicStruct{PublicField: 2}

	fmt.Println("Public Struct:", publicStruct)
}