package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)

	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing", file.Name())

	file.Close()
}

func GetFloat(fileName string) ([]float64, error) {
	file, err := OpenFile(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var nums []float64

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	CloseFile(file)

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return nums, nil
}

func main() {
	nums, err := GetFloat("13-recovering/2-file/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	fmt.Printf("Sum: %.2f\n", sum)
}