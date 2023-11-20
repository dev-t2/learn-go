// Grading Program
package main

import (
	"fmt"
	"learn-go/05-packages-c/1-grade/keyboard"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()

	checkError(err)

	if grade >= 70 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}	
}