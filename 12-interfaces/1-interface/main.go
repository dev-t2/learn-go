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

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Robot")
}

func (r Robot) Walk() {
	fmt.Println("Robot Walk")
}

func play(n NoiseMaker) {
	n.MakeSound()
	// n.Walk()
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
	play(toy)

	toy = Horn("Horn")

	toy.MakeSound()
	play(toy)

	toy = Robot("Robot")

	toy.MakeSound()
	play(toy)
}