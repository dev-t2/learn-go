package main

import "fmt"

func main() {
	var a int = 1
	var b string = "A"

	a = 2
	b = "B"

	fmt.Println(a, b)

	var c int = 3
	var d int = 4
	var e int = c * d

	fmt.Println(e)
}