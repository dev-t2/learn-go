package main

import "fmt"

type car struct {
	name 		 string
	topSpeed int
}

type part struct {
	description string
	count 			int
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

	fmt.Println()

	bolts := minimumOrder("Hex Bolts")

	showInfo(bolts)
}