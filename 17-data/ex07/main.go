package main

import "fmt"

func main() {
	var m = make(map[int]int)

	m[1] = 0
	m[2] = 2
	m[3] = 4

	delete(m, 3)
	delete(m, 4)

	var value1, isExist1 = m[3]
	var value2, isExist2 = m[4]

	fmt.Println(m[1])
	fmt.Println(value1, isExist1)
	fmt.Println(value2, isExist2)
}