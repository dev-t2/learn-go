package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE"
	utf8String := "가나다라마"

	fmt.Println(len(asciiString)) // 5
	fmt.Println(len(utf8String)) // 15

	fmt.Println()

	fmt.Println(utf8.RuneCountInString(asciiString)) // 5
	fmt.Println(utf8.RuneCountInString(utf8String)) // 5

	fmt.Println()

	asciiBytes := []byte(asciiString)
	utf8Bytes := []byte(utf8String)

	asciiBytesPartial := asciiBytes[3:]
	utf8BytesPartial := utf8Bytes[3:]

	fmt.Println(string(asciiBytesPartial)) // DE
	fmt.Println(string(utf8BytesPartial)) // 나다라마

	fmt.Println()

	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)

	asciiRunesPartial := asciiRunes[3:]
	utf8RunesPartial := utf8Runes[3:]

	fmt.Println(string(asciiRunesPartial)) // DE
	fmt.Println(string(utf8RunesPartial)) // 라마

	fmt.Println()

	for index, currentByte := range asciiBytes {
		fmt.Printf("%d: %s\n", index, string(currentByte)) // 0: A, 1: B, 2: C, 3: D, 4: E
	}

	fmt.Println()

	for index, currentByte := range utf8Bytes {
		fmt.Printf("%d: %s\n", index, string(currentByte)) // Length: 15
	}

	fmt.Println()

	for index, currentRune := range asciiRunes {
		fmt.Printf("%d: %s\n", index, string(currentRune)) // 0: A, 1: B, 2: C, 3: D, 4: E
	}

	fmt.Println()

	for index, currentRune := range utf8Runes {
		fmt.Printf("%d: %s\n", index, string(currentRune)) // 0: 가, 1: 나, 2: 다, 3: 라, 4: 마
	}
}