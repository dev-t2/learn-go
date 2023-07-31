package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var ring = ring.New(5)

	var length = ring.Len()

	for i := 0; i < length; i++ {
		ring.Value = 'A' + i

		ring = ring.Next()
	}

	for i := 0; i < length; i++ {
		fmt.Printf("%c ", ring.Value)

		ring = ring.Next()
	}

	fmt.Println()

	for i := 0; i < length; i++ {
		fmt.Printf("%c ", ring.Value)

		ring = ring.Prev()
	}
}