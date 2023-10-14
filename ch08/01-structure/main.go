package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.number, myStruct.word, myStruct.toggle)
}