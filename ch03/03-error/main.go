package main

import (
	"errors"
	"fmt"
	"math"
)

func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("Can't get square root of negative number")
	}

	return math.Sqrt(num), nil
}

func main() {
	err := errors.New("Height can't be negative")

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	err = fmt.Errorf("Height of %.2f is invalid", -2.33333)

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	root, err := squareRoot(-9.3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.3f", root)
	}
}