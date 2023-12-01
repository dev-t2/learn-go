package main

import "fmt"

func calmDown() {
	p := recover()

	err, ok := p.(error)

	if ok {
		fmt.Println(err.Error())
	}
}

func breakOut() {
	defer calmDown()

	err := fmt.Errorf("Break Out")

	panic(err)
}

func main() {
	breakOut()

	fmt.Println("Exiting")
}