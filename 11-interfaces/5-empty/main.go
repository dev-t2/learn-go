package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Make Sound")
}

func AcceptAnything(any interface{}) {
	fmt.Println(any)

	whistle, ok := any.(Whistle)

	if ok {
		whistle.MakeSound()
	}
}

func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)

	whistle.MakeSound()
}

func main() {
	AcceptAnything(3.141592)
	AcceptAnything("string")
	AcceptAnything(true)

	fmt.Println()

	AcceptAnything(Whistle("whistle"))
	AcceptWhistle(Whistle("whistle"))
}