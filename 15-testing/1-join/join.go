package join

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:len(phrases) - 1], ", ")
		result += ", and "
		result += phrases[len(phrases) - 1]
	
		return result
	}
}

func main() {
	fmt.Println(strings.Join([]string{"2023", "11", "04"}, "-"))

	fmt.Println()

	phrases := []string{"Austin", "Chloe"}

	fmt.Println(JoinWithCommas(phrases))

	phrases = []string{"Austin", "Chloe", "Sally"}

	fmt.Println(JoinWithCommas(phrases))
}