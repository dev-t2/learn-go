package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("Goodbye")

	fmt.Println("Hello")

	return fmt.Errorf("Error")
}

func main() {
	err := Socialize()

	if err != nil {
		log.Fatal(err)
	}
}