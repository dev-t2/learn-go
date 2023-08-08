package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Scan Error")
	}

	var word = scanner.Text()
	var num, err = strconv.Atoi(word)

	if err != nil {
		return 0, 0, fmt.Errorf("word: %s, err: %w", word, err)
	}

	return num, len(word), nil
}