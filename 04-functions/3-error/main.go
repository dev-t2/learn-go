package main

import (
	"errors"
	"fmt"
	"math"
)

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by 0")
	}

	return dividend / divisor, nil
}

func getSquareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("Can't square root of negative number")
	}

	return math.Sqrt(num), nil
}

func main() {
	err := errors.New("Height can't be negative")

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	err = fmt.Errorf("Height of %.2f is invalid", -1.23456)

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	result, err := divide(1.23456, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f\n", result)
	}

	result, err = getSquareRoot(-1.23456)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f\n", result)
	}
}