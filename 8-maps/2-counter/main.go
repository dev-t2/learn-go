package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getStrings(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func main() {
	lines, err := getStrings("8-maps/2-counter/votes.txt")

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