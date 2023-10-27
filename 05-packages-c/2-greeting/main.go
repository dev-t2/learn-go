package main

import (
	"fmt"
	"learn-go/05-packages-c/packages/greeting"
	"learn-go/05-packages-c/packages/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}