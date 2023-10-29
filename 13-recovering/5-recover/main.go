package main

import "fmt"

func calmDown() {
	fmt.Println("Recover:", recover())
}

func breakOut() {
	defer calmDown()

	panic("Break Out")
}

func main() {
	breakOut()

	fmt.Println("Exiting")
}