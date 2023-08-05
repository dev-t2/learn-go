package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	var file, error = os.Open(filename)

	if error != nil {
		return "", error
	}

	defer file.Close()

	var data = bufio.NewReader(file)
	var line, _ = data.ReadString('\n')

	return line, nil
}

func WriteFile(filename string, line string) error {
	var file, error = os.Create(filename)

	if error != nil {
		return error
	}

	defer file.Close()

	_, error = fmt.Fprintln(file, line)

	return error
}

const filename = "18-error/ex01/data.txt"

func main() {
	var line, error = ReadFile(filename)

	if error != nil {
		error = WriteFile(filename, "File Data")

		if error != nil {
			fmt.Println("Write File Error")

			return
		}

		line, error = ReadFile(filename)

		if error != nil {
			fmt.Println("Read File Error")

			return
		}
	}

	fmt.Println(line)
}