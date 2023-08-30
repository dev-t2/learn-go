package main

import (
	"fmt"
	"time"
)

func PrintString() {
	var str = []string{"A", "B", "C", "D", "E", "F", "G"}

	for _, s := range str {
		time.Sleep(300 * time.Millisecond)

		fmt.Printf("%s ", s)
	}
}

func PrintNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)

		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintString()
	go PrintNumber()

	time.Sleep(3 * time.Second)
}