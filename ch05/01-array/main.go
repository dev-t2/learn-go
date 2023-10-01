package main

import "fmt"

func main() {
	var nums [5]int

	nums[0] = 2
	nums[1] = 3

	fmt.Println(nums[0], nums[2], nums[4])

	fmt.Println()

	nums[0]++
	nums[0]++
	nums[2]++

	fmt.Println(nums[0], nums[2], nums[4])

	fmt.Println()

	// var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	notes := [7]string{
		"do", 
		"re", 
		"mi", 
		"fa", 
		"so", 
		"la", 
		"ti",
	}

	fmt.Println(notes[0], notes[3], notes[6])

	fmt.Println()

	fmt.Println(notes)
	fmt.Printf("%#v\n", notes)

	fmt.Println()

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
}