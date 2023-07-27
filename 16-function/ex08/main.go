package main

import "fmt"

func captureLoop1() {
	fmt.Println("Capture Loop 1")

	var f = make([]func(), 3)

	for i:=0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func captureLoop2() {
	fmt.Println("Capture Loop 2")

	var f = make([]func(), 3)

	for i := 0; i < 3; i++ {
		var v = i

		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	captureLoop1()
	captureLoop2()
}