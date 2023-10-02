package main

import (
	"fmt"
	"learn-go/ch07/02-count/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("ch07/02-count/votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}