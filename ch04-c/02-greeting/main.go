package main

import (
	"fmt"
	"learn-go/ch04-c/packages/greeting"
	"learn-go/ch04-c/packages/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}