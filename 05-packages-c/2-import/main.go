package main

import (
	"fmt"
	"learn-go/05-packages-c/2-import/greeting"
	"learn-go/05-packages-c/2-import/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}