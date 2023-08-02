package main

import "fmt"

type Product struct {
	Id    int
	Title string
}

func main() {
	var m = make(map[int]Product)

	m[16] = Product{0, "JavaScript"}
	m[46] = Product{1, "TypeScript"}
	m[78] = Product{2, "React"}
	m[345] = Product{3, "React Native"}
	m[879] = Product{4, "Go"}

	for index, value := range m {
		fmt.Println(index, value)
	}
}