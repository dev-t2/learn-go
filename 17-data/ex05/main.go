package main

import "fmt"

func main() {
	var m = make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4

	m["d"] = 5

	fmt.Println(m["a"])
	fmt.Println(m["d"])
}