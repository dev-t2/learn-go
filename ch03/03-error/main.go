package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("Height can't be negative")

	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println()

	err = fmt.Errorf("Height of %.2f is invalid", -2.33333)

	fmt.Println(err)
	log.Fatal(err)
}