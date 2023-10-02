package main

import (
	"fmt"
	"learn-go/packages/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("ch07/01-count/votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)
}