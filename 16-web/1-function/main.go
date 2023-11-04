package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(function func()) {
	function()
	function()
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}

func multiply(a, b int) float64 {
	return float64(a) * float64(b)
}

func doMath(function func(int, int) float64) {
	result := function(10, 2)

	fmt.Println(result)
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

	fmt.Println()

	doMath(divide)
	doMath(multiply)
}