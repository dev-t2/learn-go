package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func getDouble(num float64) float64 {
	// return int(num * 2)

	return num * 2
}

func main() {
	sayHi()

	fmt.Println()

	repeatLine("hello", 3)

	fmt.Println()

	fmt.Println(getDouble(6.0)) 
}