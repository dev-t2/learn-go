package main

import "fmt"

func AcceptAnything(any interface{}) {
	fmt.Println(any)
}

func main() {
	AcceptAnything(3.141592)
	AcceptAnything("string")
	AcceptAnything(true)
}