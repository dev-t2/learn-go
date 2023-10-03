package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part

	p.description = description
	p.count = 100

	return p
}

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

	fmt.Println()

	var ev6 car

	ev6.name = "KIA EV6 GT"
	ev6.topSpeed = 260

	fmt.Println("Name:", ev6.name)
	fmt.Println("Top Speed:", ev6.topSpeed)

	fmt.Println()

	var bolts part

	bolts.description = "Bolts"
	bolts.count = 25

	showInfo(bolts)

	fmt.Println()

	p := minimumOrder("Bolts")

	fmt.Println(p.description, p.count)
}