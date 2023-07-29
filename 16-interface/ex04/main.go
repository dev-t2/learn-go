package main

import "fmt"

type Stringer interface {
	String() string
}

type User struct {
	Id int
}

func (u *User) String() string {
	return fmt.Sprintf("User Id: %d", u.Id)
}

func PrintId(stringer Stringer) {
	var s = stringer.(*User)

	fmt.Printf("Id: %d\n", s.Id)
}

func main() {
	var user = &User{1}

	PrintId(user)
}