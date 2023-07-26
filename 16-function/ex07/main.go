package main

import "fmt"

func main() {
	var num = 0

	var function = func() {
		num += 10
	}

	num++

	function()

	fmt.Println(num)
}