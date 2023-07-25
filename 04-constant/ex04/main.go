package main

import "fmt"

const (
	A1 int = iota
	A2 int = iota
	A3 int = iota	
)

const (
	B1 = iota
	B2
	B3
)

const (
	C1 = iota + 1
	C2
	C3
)

func main() {
	fmt.Println(A3, B3, C3)
}