package main

import "fmt"

type car struct {
	name     string
	topSpeed int
}

func main() {
	var myStruct struct {
		num    float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	myStruct.num = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.num, myStruct.word, myStruct.toggle)

	fmt.Println()

	var ev6 car

	ev6.name = "KIA EV6 GT"
	ev6.topSpeed = 260

	fmt.Println("Name:", ev6.name)
	fmt.Println("Top Speed:", ev6.topSpeed)
}