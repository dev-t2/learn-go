package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	for i := 3; i >= 1; i-- {
		fmt.Println(i)
	}

	fmt.Println()

	for i := 1; i <= 5; i += 2 {
		fmt.Println(i)
	}

	fmt.Println()

	for i := 15; i >= 5; i -= 5 {
		fmt.Println(i)
	}

	fmt.Println()

	i := 1

	for i <= 3 {
		fmt.Println(i)

		i++
	}

	fmt.Println()

	i = 3

	for i >= 1 {
		fmt.Println(i)

		i--
	}

	fmt.Println()

	for x := 1; x <= 3; x++ {
		y := x + 1

		fmt.Println(y)
	}

	// fmt.Println(y)

	fmt.Println()

	for x := 1; x <= 3; x++ {
		fmt.Println(x)
	}

	// fmt.Println(x)

	fmt.Println()

	var x int

	for x = 1; x <= 3; x++ {
		fmt.Println(x)
	}

	fmt.Println(x)

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("before continue")

		continue

		// fmt.Println("after continue")
	}

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("before break")

		break

		// fmt.Println("after break")
	}

	fmt.Println("after loop")
}