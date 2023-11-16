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

	for i := 5; i >= 1; i -= 2 {
		fmt.Println(i)
	}

	fmt.Println()

	i := 1

	for i <= 3 {
		fmt.Println(i)

		i++
	}

	fmt.Println(i)

	fmt.Println()

	for i > 1 {
		fmt.Println(i)

		i--
	}

	fmt.Println(i)

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("Before Continue")

		continue

		// fmt.Println("After Continue")
	}

	fmt.Println("After Continue")

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("Before Break")

		break

		// fmt.Println("After Break")
	}

	fmt.Println("After Break")
}