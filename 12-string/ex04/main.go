package main

import "fmt"

func main() {
	var str1 = "Hello"
	var str2 = "World"
	var str3 = str1 + " " + str2

	fmt.Println(str3)

	str1 += " " + str2

	fmt.Println(str1)

	fmt.Println()

	str1 = "Hello"
	str2 = "Hell"
	str3 = "Hello"

	fmt.Printf("%s == %s: %v\n", str1, str2, str1 == str2)
	fmt.Printf("%s != %s: %v\n", str1, str2, str1 != str2)
	fmt.Printf("%s == %s: %v\n", str1, str3, str1 == str3)
	fmt.Printf("%s != %s: %v\n", str1, str3, str1 != str3)

	fmt.Println()

	str1 = "BBB"
	str2 = "aaaaAAA"
	str3 = "BBAD"

	fmt.Printf("%s > %s: %v\n", str1, str2, str1 > str2)
	fmt.Printf("%s < %s: %v\n", str1, str3, str1 < str3)
}