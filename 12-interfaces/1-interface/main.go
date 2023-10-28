package main

import (
	"fmt"
	"learn-go/12-interfaces/mypackage"
)

type NoiseMaker interface {
	MakeSound()
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Whistle")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Horn")
}

func main() {
	var value mypackage.MyInterface

	value = mypackage.MyInt(5)

	value.MethodWithoutParameter()
	value.MethodWithParameter(123.4)
	fmt.Println(value.MethodWithReturn())

	fmt.Println()

	var toy NoiseMaker

	toy = Whistle("Whistle")

	toy.MakeSound()

	toy = Horn("Horn")

	toy.MakeSound()
}