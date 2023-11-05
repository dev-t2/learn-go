package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("15-web/6-file/data.txt", os.O_WRONLY, os.FileMode(0600))

	check(err)

	_, err = file.Write([]byte("Chloe\n"))

	check(err)

	err = file.Close()

	check(err)
}