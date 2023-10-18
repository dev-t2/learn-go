package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi")
}

func main() {
	value := MyType("Value")

	value.sayHi()

	anotherValue := MyType("Another Value")

	anotherValue.sayHi()
}