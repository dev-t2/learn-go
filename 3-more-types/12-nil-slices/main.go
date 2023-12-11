package main

import "fmt"

func main() {
	var s []int

	fmt.Println(s == nil)
	fmt.Println(s, len(s), cap(s))
}