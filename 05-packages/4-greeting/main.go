package main

import (
	"fmt"
	greeting "learn-go/05-packages/3-greeting"
	"learn-go/05-packages/3-greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}