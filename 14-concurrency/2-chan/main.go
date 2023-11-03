package main

import (
	"fmt"
	"time"
)

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func report(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "stop")

		time.Sleep(time.Second)
	}

	fmt.Println(name, "start")
}

func send(channel chan string) {
	report("send", 2)

	fmt.Println("*** value ***")

	channel <- "a"

	fmt.Println("*** value ***")

	channel <- "b"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)

	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)

	fmt.Println()

	channel3 := make(chan string)

	go send(channel3)

	report("receive", 5)

	fmt.Println(<-channel3)
	fmt.Println(<-channel3)
}