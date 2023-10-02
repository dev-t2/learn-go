package main

import "fmt"

func main() {
	var myStruct struct {
		num    float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	fmt.Println()

	myStruct.num = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.num, myStruct.word, myStruct.toggle)
}