package main

import "fmt"

type car struct {
	name     string
	topSpeed int
}

type part struct {
	description string
	count       int
}

type myStruct2 struct {
	myField int
}

func minimumOrder(description string) part {
	var p part

	p.description = description
	p.count = 100

	return p
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func main() {
	var myStruct1 struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct1)
	fmt.Println(myStruct1.number, myStruct1.word, myStruct1.toggle)

	fmt.Println()

	myStruct1.number = 3.14
	myStruct1.word = "pie"
	myStruct1.toggle = true

	fmt.Printf("%#v\n", myStruct1)
	fmt.Println(myStruct1.number, myStruct1.word, myStruct1.toggle)

	fmt.Println()

	var porsche car

	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)

	fmt.Println()

	bolts := minimumOrder("Hex Bolts")

	showInfo(bolts)

	fmt.Println()

	value1 := 2
	pointer1 := &value1

	fmt.Println(pointer1)
	fmt.Println(*pointer1)

	fmt.Println()

	var value2 myStruct2

	value2.myField = 3
	pointer2 := &value2

	fmt.Println((*pointer2).myField)
	fmt.Println(pointer2.myField)

	fmt.Println()
	
	(*pointer2).myField = 6

	fmt.Println(pointer2.myField)

	pointer2.myField = 9

	fmt.Println(pointer2.myField)
}