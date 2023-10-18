package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println(m)
}

func (m MyType) MethodWithParameters(num int, flag bool) {
	fmt.Println(m, num, flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func (m MyType) ExportedMethod() {}

func (m MyType) unexportedMethod() {}

func main() {
	value := MyType("Value")

	value.sayHi()

	anotherValue := MyType("Another Value")

	anotherValue.sayHi()

	fmt.Println()

	value.MethodWithParameters(4, true)
	fmt.Println(value.WithReturn())
}