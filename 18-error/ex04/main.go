package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Scan Error")
	}

	var word = scanner.Text()

	num, err := strconv.Atoi(word)

	if err != nil {
		return 0, 0, fmt.Errorf("Word: %s, Error: %w", word, err)
	}

	return num, len(word), nil
}

func MultipleFromString(str string) (int, error) {
	var scanner = bufio.NewScanner(strings.NewReader(str))

	scanner.Split(bufio.ScanWords)

	var pos = 0

	a, n, err := readNextInt(scanner)

	if err != nil {
		return 0, fmt.Errorf("Position: %d, Error: %w", pos, err)
	}

	pos += n + 1

	b, n, err := readNextInt(scanner)

	if err != nil {
		return 0, fmt.Errorf("Position: %d, Error: %w", pos, err)
	}

	return a * b, nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)

	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)

		var numError *strconv.NumError

		if errors.As(err, &numError) {
			fmt.Println(numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}