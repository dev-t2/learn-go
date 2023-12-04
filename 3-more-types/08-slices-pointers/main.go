package main

import "fmt"

func main() {
	names := [4]string{"HTML", "CSS", "JavaScript", "Go"}

	fmt.Println(names)

	fmt.Println()

	a := names[1:3]
	b := names[2:4]

	fmt.Println(a, b)

	fmt.Println()

	b[0] = "TypeScript"

	fmt.Println(a, b)
	fmt.Println(names)
}