package main

import (
	"fmt"
	"os"
)

func main() {
	var file, error = os.Create("04-function/ex04/main.txt")

	if error != nil {
		fmt.Println("Create File Error")

		return
	}

	defer fmt.Println("Call Defer")

	defer file.Close()

	defer fmt.Println("File Close")

	fmt.Println("Create File")

	fmt.Fprintln(file, "Hello World")
}