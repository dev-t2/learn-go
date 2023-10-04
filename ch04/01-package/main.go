package main

import (
	"fmt"
	"learn-go/packages/greeting"
	"learn-go/packages/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}