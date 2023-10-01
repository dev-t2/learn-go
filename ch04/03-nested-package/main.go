package main

import (
	"fmt"
	"learn-go/ch04/03-nested-package/greeting"
	"learn-go/ch04/03-nested-package/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	deutsch.Hallo()
	deutsch.GutenTag()
}