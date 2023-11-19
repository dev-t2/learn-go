package main

import "fmt"

func printLine() {
	fmt.Println("hello")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func returns() (string, int, bool) {
	return "hello", 1, true
}

func main() {
	printLine()

	fmt.Println()

	repeatLine("hello", 3)

	fmt.Println()

	myInt, myBool, myString := returns()

	fmt.Println(myInt, myBool, myString)
}