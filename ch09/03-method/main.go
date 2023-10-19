package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println(m)
}

func (m MyType) MethodWithParameters(num int) {
	fmt.Println(m, num)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func (m MyType) ExportedMethod() {}

func (m MyType) unexportedMethod() {}

func main() {
	value1 := MyType("Value1")

	value1.sayHi()

	fmt.Println()

	value2 := MyType("Value2")

	value2.MethodWithParameters(4)

	fmt.Println(value2.WithReturn())
}