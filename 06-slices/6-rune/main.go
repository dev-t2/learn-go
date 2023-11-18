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

	fmt.Println()

	asciiBytes := []byte(asciiString)
	utf8Bytes := []byte(utf8String)

	asciiBytesPartial := asciiBytes[3:]
	utf8BytesPartial := utf8Bytes[3:]

	fmt.Println(string(asciiBytesPartial))
	fmt.Println(string(utf8BytesPartial))

	fmt.Println()

	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)

	asciiRunesPartial := asciiRunes[3:]
	utf8RunesPartial := utf8Runes[3:]

	fmt.Println(string(asciiRunesPartial))
	fmt.Println(string(utf8RunesPartial))

	fmt.Println()

	for index, currentByte := range asciiBytes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	fmt.Println()

	for index, currentByte := range utf8Bytes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	fmt.Println()

	for index, currentRune := range asciiRunes {
		fmt.Printf("%d: %s\n", index, string(currentRune))
	}

	fmt.Println()

	for index, currentRune := range utf8Runes {
		fmt.Printf("%d: %s\n", index, string(currentRune))
	}
}