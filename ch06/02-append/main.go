package main

import "fmt"

func main() {
	slice := []string{"a", "b"}

	fmt.Println(slice, len(slice))

	slice = append(slice, "c")

	fmt.Println(slice, len(slice))

	slice = append(slice, "d", "e")

	fmt.Println(slice, len(slice))
}