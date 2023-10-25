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
	value := MyType("Value")

	value.sayHi()

	value.MethodWithParameters(5)

	fmt.Println(value.WithReturn())

	fmt.Println()

	value.method()
	(&value).pointerMethod()
	value.pointerMethod()

	fmt.Println()

	pointer := &value

	(*pointer).method()
	pointer.method()
	pointer.pointerMethod()

	// &MyType("Value")
	// MyType("Value").method()
	// MyType("Value").pointerMethod()

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