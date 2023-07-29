package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func HelloWorld(writer Writer) {
	writer("Hello World")
}

func main() {
	var file, error = os.Create("04-function/ex09/main.txt")

	if error != nil {
		fmt.Println("Create File Error")

		return
	}

	defer file.Close()

	HelloWorld(func (message string) {
		fmt.Fprintln(file, message)
	})
}