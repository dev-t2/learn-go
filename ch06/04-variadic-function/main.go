package main

import "fmt"

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()

	fmt.Println()

	mix(1, true, "a", "b")
}