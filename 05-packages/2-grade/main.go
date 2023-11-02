// Grading Program
package main

import (
	"fmt"
	keyboard "learn-go/05-packages/1-keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 70 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}	
}