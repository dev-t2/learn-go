package main

import (
	"fmt"
	"math"
)

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func returns() (int, bool, string) {
	return 1, true, "Go"
}

func floatParts(num float64) (integerPart int, floatPart float64) {
	wholeNumber := math.Floor(num)

	return int(wholeNumber), num - wholeNumber
}

func main() {
	sayHi()

	fmt.Println()

	repeatLine("hello", 3)

	fmt.Println()

	myInt, myBool, myString := returns()

	fmt.Println(myInt, myBool, myString)
	fmt.Println(returns())

	fmt.Println()
	
	fmt.Println(floatParts(1.26))
}