package main

import "fmt"

const (
	A1 = iota
	A2
	A3
)

const (
	B1 = iota + 1
	B2
	B3
)

func main() {
	fmt.Println(A3, B3)
}