package main

import "fmt"

func main() {
	var char = '용'

	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)

	var str1 = "가나다라마"
	var str2 = "abcde"

	fmt.Printf("len(str1) = %d\n", len(str1))
	fmt.Printf("len(str2) = %d\n", len(str2))

	var str3 = "Hello World"
	var runes1 = []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}


	fmt.Println(str3)
	fmt.Println(string(runes1))

	var str4 = "Hello 월드"
	var runes2 = []rune(str4)

	fmt.Printf("len(str4) = %d\n", len(str4))
	fmt.Printf("len(runes2) = %d\n", len(runes2))
}