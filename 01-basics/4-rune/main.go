package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE"
	utf8String := "가나다라마"

	fmt.Println(len(asciiString))
	fmt.Println(len(utf8String))

	fmt.Println()

	fmt.Println(utf8.RuneCountInString(asciiString))
	fmt.Println(utf8.RuneCountInString(utf8String))
}