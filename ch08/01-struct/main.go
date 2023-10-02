package main

import "fmt"

func main() {
	var myStruct struct {
		num    float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)
}