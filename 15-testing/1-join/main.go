package main

import (
	"fmt"
	"strings"
)

func joinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases) - 1], ", ")
	result += ", and "
	result += phrases[len(phrases) - 1]

	return result
}

func main() {
	fmt.Println(strings.Join([]string{"2023", "11", "04"}, "-"))

	fmt.Println()

	phrases := []string{"Austin", "Chloe"}

	fmt.Println(joinWithCommas(phrases))

	phrases = []string{"Austin", "Chloe", "Sally"}

	fmt.Println(joinWithCommas(phrases))
}