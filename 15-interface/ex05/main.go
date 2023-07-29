package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct{}

func (f *File) Read() {}

func ReadFile(reader Reader) {
	fmt.Println("Read File Start")

	if file, isSuccess := reader.(Closer); isSuccess {
		fmt.Println("File Close")

		file.Close()
	}

	fmt.Println("Read File End")
}

func main() {
	var file = &File{}

	ReadFile(file)
}