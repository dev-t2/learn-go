package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id   	 int32
	Height float64
}

func PrintUser(user User) {
	fmt.Printf("아이디: %d, 이름: %.1f\n", user.Id, user.Height)
}

func main() {
	var user1 = User{1, 173.0}
	var user2 = user1

	PrintUser(user2)

	fmt.Println(unsafe.Sizeof(user2))
}