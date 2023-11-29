package main

import "fmt"

func main() {
	nums := make([]int, 5)

	nums[2] = 2
	nums[4] = 4

	fmt.Println(nums[0], nums[2], nums[4])

	nums[0]++
	nums[0]++
	nums[2]++

	fmt.Println(nums[0], nums[2], nums[4])

	fmt.Println()

	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}

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

	fmt.Println()

	array := [5]string{"a", "b", "c", "d", "e"}

	slice1 := array[0:3]

	fmt.Println(slice1)

	slice2 := array[2:5]

	fmt.Println(slice2)

	slice3 := array[:3]

	fmt.Println(slice3)

	slice4 := array[1:]

	fmt.Println(slice4)

	fmt.Println()

	array[1] = "x"

	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}