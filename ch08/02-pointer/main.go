package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	value1 := 2
	pointer1 := &value1

	fmt.Println(pointer1)
	fmt.Println(*pointer1)

	fmt.Println()

	var value2 myStruct

	value2.myField = 3
	pointer2 := &value2

	fmt.Println((*pointer2).myField)
	fmt.Println(pointer2.myField)
	
	(*pointer2).myField = 6

	fmt.Println(pointer2.myField)

	pointer2.myField = 6

	fmt.Println(pointer2.myField)
}