// Grading Program
package main

import (
	"fmt"
	"learn-go/05-packages/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}	
}