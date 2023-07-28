package main

import "fmt"

func main() {
	var n1 = 10
	var n2 = 20

	var p1 = &n1
	var p2 = &n1
	var p3 = &n2

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)
}