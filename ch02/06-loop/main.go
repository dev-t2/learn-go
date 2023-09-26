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

	for i >= 1 {
		fmt.Println(i)

		i--
	}

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println(i)

		j := i + 1

		fmt.Println(j)
	}

	// fmt.Println(i)
	// fmt.Println(j)

	fmt.Println()

	var x int

	for x = 1; x <= 3; x++ {
		fmt.Println(x)
	}

	fmt.Println(x)

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("Before Continue")

		continue

		// fmt.Println("After Continue")
	}

	fmt.Println("After Loop")

	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println("Before Break")

		break

		// fmt.Println("After Break")
	}

	fmt.Println("After Loop")
}