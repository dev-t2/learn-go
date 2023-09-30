// Celsius Program
package main

import (
	"fmt"
	"learn-go/ch04/01-package/keyboard"
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