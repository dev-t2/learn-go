package main

import "fmt"

type car struct {
	name 		 string
	topSpeed int
}

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

	fmt.Println()

	var porsche car

	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)
}