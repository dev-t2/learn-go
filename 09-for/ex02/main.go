package main

import (
	"fmt"
	"time"
)

func main() {
	var i = 0

	for {
		time.Sleep(time.Second)

		fmt.Println(i)

		i++
	}
}