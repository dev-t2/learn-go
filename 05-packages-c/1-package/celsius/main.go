// Celsius Program
package main

import (
	"fmt"
	"learn-go/05-packages-c/packages/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")

	fahrenheit, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.2f degrees celsius\n", celsius)
}