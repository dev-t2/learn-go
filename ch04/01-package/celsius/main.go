// Celsius Program
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
	fmt.Print("Enter a temperature in Fahrenheit: ")

	fahrenheit, err := getFloat()

	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.2f degrees celsius\n", celsius)
}