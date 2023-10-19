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

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

type Number int

func (n *Number) double() {
	*n *= 2
}

func main() {
	value1 := MyType("Value1")

	value1.sayHi()

	fmt.Println()

	value2 := MyType("Value2")

	value2.MethodWithParameters(4)

	fmt.Println(value2.WithReturn())

	fmt.Println()

	value3 := MyType("Value 3")

	value3.method()
	(&value3).pointerMethod()
	value3.pointerMethod()

	fmt.Println()

	pointer := &value3

	(*pointer).method()
	pointer.method()
	pointer.pointerMethod()

	fmt.Println()

	num := Number(4)
	
	fmt.Println(num)

	num.double()

	fmt.Println(num)
}