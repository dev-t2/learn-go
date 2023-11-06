package main

import "fmt"

func main() {
	fmt.Printf("0 & 0 == %b\n", 0 & 0)
	fmt.Printf("1 & 0 == %b\n", 1 & 0)
	fmt.Printf("1 & 1 == %b\n", 1 & 1)

	fmt.Println()

	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 1 & 3)

	fmt.Println()

	fmt.Printf("0 | 0 == %b\n", 0 | 0)
	fmt.Printf("1 | 0 == %b\n", 1 | 0)
	fmt.Printf("1 | 1 == %b\n", 1 | 1)

	fmt.Println()

	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 1 | 3)
}