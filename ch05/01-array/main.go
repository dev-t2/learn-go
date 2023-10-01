package main

import "fmt"

func main() {
	var nums [5]int

	nums[0] = 2
	nums[1] = 3

	fmt.Println(nums[0])
	fmt.Println(nums[2])
	fmt.Println(nums[4])

	fmt.Println()

	nums[0]++
	nums[0]++
	nums[2]++

	fmt.Println(nums[0])
	fmt.Println(nums[2])
	fmt.Println(nums[4])
}