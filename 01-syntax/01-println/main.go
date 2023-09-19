package main

import (
	"fmt"
	"math"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("First Argument", "Second Argument")

	fmt.Println()

	fmt.Println(math.Floor(2.75))
	fmt.Println(cases.Title(language.English).String("go language"))

	fmt.Println()

	fmt.Println("Hello,\nGo!")
	fmt.Println("Hello,\tGo!")
	fmt.Println("Quotes: \"\"")
	fmt.Println("Backslash: \\")
}