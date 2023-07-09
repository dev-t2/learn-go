package main

import "fmt"

func main() {
	var n int = 100
	var p *int

	p = &n

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 200

	fmt.Printf("n의 값: %d\n", n)
}