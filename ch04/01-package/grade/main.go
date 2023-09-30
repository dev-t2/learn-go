// Grading Program
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	return num, nil
}

func main() {
	fmt.Print("Enter a grade: ")

	grade, err := getFloat()

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}	
}