package main

import "fmt"

func main() {
	var nums = [...]int{1, 2, 3, 4, 5}

	nums[2] = 30

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}