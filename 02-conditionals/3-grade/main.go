package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	str = strings.TrimSpace(str)
	grade, err := strconv.ParseFloat(str, 64)

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}	
}