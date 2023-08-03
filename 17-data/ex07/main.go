package main

import "fmt"

func main() {
	var m = make(map[int]int)

	m[1] = 0
	m[2] = 2
	m[3] = 4

	delete(m, 3)
	delete(m, 4)

	fmt.Println(m[3])
	fmt.Println(m[1])
}