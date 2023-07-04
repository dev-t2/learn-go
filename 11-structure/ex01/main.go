package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	var user User

	user.Id = 0
	user.Name = "TypeScript"

	fmt.Println("id:", user.Id)
	fmt.Println("name:", user.Name)
}