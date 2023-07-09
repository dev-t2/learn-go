package main

import "fmt"

func main() {
	var str1 = "Hello\tWorld\n"

	fmt.Println(str1)

	var str2 = `Hello\tWorld\n`

	fmt.Println(str2)

	fmt.Println()

	var str3 = "Hello\nWorld\n"

	fmt.Println(str3)

	var str4 = `Hello
World`

	fmt.Println(str4)
}