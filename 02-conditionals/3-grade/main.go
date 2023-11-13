package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
} 

func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	checkError(err)

	str = strings.TrimSpace(str)
	grade, err := strconv.ParseFloat(str, 64)

	checkError(err)

	if grade >= 70 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}	
}