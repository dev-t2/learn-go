package main

import "fmt"

type customInt int

func (a customInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a customInt = 10

	fmt.Println(a.add(30))

	var b = 20

	fmt.Println(customInt(b).add(50))
}