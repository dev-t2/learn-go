package custom

import "fmt"

var a = 3

func init() {
	a++

	fmt.Println("Init Function - a:", a)
}

func customFunction() int {
	a++

	fmt.Println("Custom Function - a:", a)

	return a
}

var b = customFunction()
var c = customFunction()
var d = b + c

func CustomPrint() {
	fmt.Println("Custom Print Function - a:", a)
}