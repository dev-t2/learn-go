package main

import "fmt"

func count(start, end int) {
	fmt.Printf("count(%d, %d)\n", start, end)
	
	if start < end {
		count(start + 1, end)
	}

	fmt.Printf("Returning from count(%d, %d)\n", start, end)
}

func main() {
	count(1, 3)
}