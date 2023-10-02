package main

import "fmt"

func main() {
	// var ranks map[string]int
	// ranks = make(map[string]int)

	ranks := make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks["gold"])
	fmt.Println(ranks["silver"])
	fmt.Println(ranks["bronze"])

	fmt.Println()

	ranks = map[string]int{"gold": 4, "silver": 5, "bronze": 6}

	fmt.Println(ranks["gold"])
	fmt.Println(ranks["silver"])
	fmt.Println(ranks["bronze"])

	fmt.Println()

	emptyMap := map[string]int{}

	fmt.Println(emptyMap)
}