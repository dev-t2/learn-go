package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var ev6 car

	ev6.name = "KIA EV6 GT"
	ev6.topSpeed = 260

	fmt.Println("Name:", ev6.name)
	fmt.Println("Top Speed:", ev6.topSpeed)
}