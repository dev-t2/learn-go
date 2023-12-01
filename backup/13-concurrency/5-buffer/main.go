package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan string) {
	time.Sleep(1 * time.Second)

	channel<-"a"

	time.Sleep(1 * time.Second)

	channel<-"b"

	time.Sleep(1 * time.Second)

	channel<-"c"
}

func main() {
	channel := make(chan string, 3)

	fmt.Println(time.Now())

	fmt.Println()

	go sendLetters(channel)

	time.Sleep(5 * time.Second)

	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())

	fmt.Println()

	fmt.Println(time.Now())
}