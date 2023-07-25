package main

import (
	"fmt"
	"unsafe"
)

type User1 struct {
	A int8
	B int
	C int8
	D int
	E int8
}

type User2 struct {
	A int8
	C int8
	E int8
	B int
	D int
}

func main() {
	var user1 = User1{1, 2, 3, 4, 5}

	fmt.Println(unsafe.Sizeof(user1))

	var user2 = User2{1, 2, 3, 4, 5}

	fmt.Println(unsafe.Sizeof(user2))
}