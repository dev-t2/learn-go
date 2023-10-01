package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([3]float64, error) {
	var nums [3]float64

	file, err := os.Open(filename)

	if err != nil {
		return nums, err
	}

	i := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums[i], err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nums, err
		}

		i++
	}

	err = file.Close()

	if err != nil {
		return nums, err
	}

	return nums, nil
}