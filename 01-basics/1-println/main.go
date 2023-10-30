package main

import (
	"fmt"
	"math"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Argument1", "Argument2")

	fmt.Println()

	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Ceil(1.5))

	fmt.Println()
	
	fmt.Println(cases.Title(language.English).String("go language"))

	fmt.Println()

	fmt.Println("\"Escape Sequence\"")
	fmt.Println("Escape\nSequence")
}