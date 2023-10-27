package main

import "fmt"

func main() {
	slice := []string{"a", "b"}

	fmt.Println(slice, len(slice))

	slice = append(slice, "c")

	fmt.Println(slice, len(slice))

	slice = append(slice, "d", "e")

	fmt.Println(slice, len(slice))

	fmt.Println()

	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")

	fmt.Println(s1, s2, s3, s4)

	s4[0] = "xx"

	fmt.Println(s1, s2, s3, s4)

	fmt.Println()

	s1 = []string{"s1", "s1"}
	s1 = append(s1, "s2", "s2")
	s1 = append(s1, "s3", "s3")
	s1 = append(s1, "s4", "s4")

	fmt.Println(s1)

	fmt.Println()

	var intSlice []int

	fmt.Printf("intSlice: %#v\n", intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(intSlice == nil)

	fmt.Println()

	intSlice = append(intSlice, 27)

	fmt.Printf("intSlice: %#v\n", intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(intSlice == nil)
}