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

	counter := make(map[string]int)

	counter["a"]++
	counter["a"]++
	counter["c"]++

	fmt.Println(counter["a"], counter["b"], counter["c"])

	fmt.Println()

	var nilMap map[int]string

	fmt.Printf("%#v\n", nilMap)

	// nilMap[3] = "three"

	nilMap = make(map[int]string)

	nilMap[3] = "three"

	fmt.Printf("%#v\n", nilMap)

	fmt.Println()

	value, ok := counter["a"]

	fmt.Println(value, ok)

	value, ok = counter["b"]

	fmt.Println(value, ok)

	value, ok = counter["c"]

	fmt.Println(value, ok)

	fmt.Println()

	delete(counter, "a")

	value, ok = counter["a"]

	fmt.Println(value, ok)
}