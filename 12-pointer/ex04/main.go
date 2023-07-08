package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func CreateUser(id int, name string) *User {
	var u = User{id, name}

	return &u
}

func main() {
	var user = CreateUser(1, "Go")

	fmt.Println(user)
}