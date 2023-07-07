package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func PrintUser(user User) {
	fmt.Printf("아이디: %d, 이름: %s", user.Id, user.Name)
}

func main() {
	var user1 = User{1, "Go"}
	var user2 = user1

	PrintUser(user2)
}