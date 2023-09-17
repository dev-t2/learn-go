package main

import (
	"fmt"
	"math"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(cases.Title(language.English).String("go language"))
}