package main

import (
	"fmt"
	"log"
)

type CustomError string

func (c CustomError) Error() string {
	return string(c)
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %.2f degrees", o)
}

func checkTemperature(actual, safe float64) error {
	excess := actual - safe

	if excess > 0 {
		return OverheatError(excess)
	}

	return nil
}

func main() {
	var err error

	err = CustomError("Custom Error")

	fmt.Println(err)

	fmt.Println()

	err = checkTemperature(123.456, 100.0)

	if err != nil {
		log.Fatal(err)
	}
}