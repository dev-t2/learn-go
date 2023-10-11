package main

import (
	"fmt"
	"sort"
)

func main() {
	// var ranks map[string]int
	// ranks = make(map[string]int)

	ranks := make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks["gold"], ranks["silver"], ranks["bronze"])

	ranks = map[string]int{"gold": 4, "silver": 5, "bronze": 6}

	fmt.Println(ranks["gold"], ranks["silver"], ranks["bronze"])

	fmt.Println()

	counter := make(map[string]int)

	counter["a"]++
	counter["a"]++
	counter["c"]++

	fmt.Println(counter["a"], counter["b"], counter["c"])

	fmt.Println()

	value, ok := counter["a"]

	fmt.Println(value, ok)

	value, ok = counter["b"]

	fmt.Println(value, ok)

	value, ok = counter["c"]

	fmt.Println(value, ok)

	delete(counter, "a")

	value, ok = counter["a"]

	fmt.Println(value, ok)

	fmt.Println()

	counter = map[string]int{"a": 1, "b": 2, "c": 3}

	var keys []string

	for key, value := range counter {
		fmt.Println(key, value)

		keys = append(keys, key)
	}

	fmt.Println()

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, counter[key])
	}

	fmt.Println()

	var nilMap map[int]string

	fmt.Printf("%#v\n", nilMap)
	fmt.Println(nilMap == nil)

	// nilMap[3] = "three"

	fmt.Println()

	nilMap = make(map[int]string)

	nilMap[3] = "three"

	fmt.Printf("%#v\n", nilMap)
	fmt.Println(nilMap == nil)
}