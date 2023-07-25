package main

import "fmt"

type Stringer interface {
	String() string
}

type User struct {
	Id int
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d", u.Id)
}

func main() {
	var user = User{1}
	var stringer Stringer

	stringer = user

	fmt.Printf("%s\n", stringer.String())
}