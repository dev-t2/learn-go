package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice = []int{5, 2, 6, 3, 1, 4}

	fmt.Println(slice)

	sort.Ints(slice)

	fmt.Println(slice)
}