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

	fmt.Println(lines)
}