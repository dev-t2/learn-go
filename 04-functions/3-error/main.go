package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by 0")
	}

	return a / b, nil
}

func main() {
	err := errors.New("Height can't be negative")

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	err = fmt.Errorf("Height of %.2f is invalid", -1.23456789)

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	result, err := divide(15, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f\n", result)
	}
}