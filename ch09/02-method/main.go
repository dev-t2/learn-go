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

type Gallons float64

type Liters float64

type Milliliters float64

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) toGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.787879)
}

func (g Gallons) toMilliliters() Milliliters {
	return Milliliters(g * 3787.879)
}

func main() {
	value1 := MyType("Value1")

	value1.sayHi()

	value1.MethodWithParameters(4)

	fmt.Println(value1.WithReturn())

	fmt.Println()

	value2 := MyType("Value2")

	value2.method()
	(&value2).pointerMethod()
	value2.pointerMethod()

	fmt.Println()

	pointer := &value2

	(*pointer).method()
	pointer.method()
	pointer.pointerMethod()

	// &MyType("Value4")
	// MyType("Value4").method()
	// MyType("Value4").pointerMethod()

	fmt.Println()

	num := Number(4)
	
	fmt.Println(num)

	num.double()

	fmt.Println(num)

	fmt.Println()

	water1 := Liters(2)

	fmt.Printf("%.1f liters equals %.3f gallons\n", water1, water1.toGallons())

	water2 := Milliliters(2000)

	fmt.Printf("%.1f milliliters equals %.3f gallons\n", water2, water2.toGallons())

	fmt.Println()

	water3 := Gallons(0.528)

	fmt.Printf("%.3f gallons equals %.1f liters\n", water3, water3.toLiters())
	fmt.Printf("%.3f gallons equals %.1f milliliters\n", water3, water3.toMilliliters())
}