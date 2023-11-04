package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	// var myFunction func()
	// myFunction = sayHi

	// var myFunction func() = sayHi

	myFunction := sayHi

	myFunction()

	fmt.Println()

	twice(sayBye)

	fmt.Println()

	var mathFunction func(int, int) float64
	
	mathFunction = divide

	fmt.Println(mathFunction(4, 2))
}