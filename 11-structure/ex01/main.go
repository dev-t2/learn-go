package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	var user User

	user.Id = 0
	user.Name = "Hello"

	fmt.Println("id:", user.Id)
	fmt.Println("name:", user.Name)

	user = User {1, "World"}

	fmt.Println("id:", user.Id)
	fmt.Println("name:", user.Name)

	user = User {Name: "Hello World"}

	fmt.Println("id:", user.Id)
	fmt.Println("name:", user.Name)
}