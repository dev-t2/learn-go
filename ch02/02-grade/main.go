package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	
	log.Fatal(err)

	fmt.Print(input)
}