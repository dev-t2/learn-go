package main

import "fmt"

func main() {
	var nums [5]int

	nums[2] = 2
	nums[4] = 4

	fmt.Println(nums[0], nums[2], nums[4])

	nums[0]++
	nums[0]++
	nums[2]++

	fmt.Println(nums[0], nums[2], nums[4])

	fmt.Println()

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	fmt.Println(notes[0], notes[3], notes[6])
	fmt.Println(notes)
	fmt.Printf("%#v\n", notes)

	fmt.Println()

	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	fmt.Println()

	for _, note := range notes {
		fmt.Println(note)
	}
}