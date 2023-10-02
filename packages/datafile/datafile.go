package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	var nums []float64

	file, err := os.Open(filename)

	if err != nil {
		return nums, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nums, err
		}

		nums = append(nums, num)
	}

	err = file.Close()

	if err != nil {
		return nums, err
	}

	if scanner.Err() != nil {
		return nums, scanner.Err()
	}

	return nums, nil
}