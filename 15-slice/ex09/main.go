package main

import (
	"fmt"
	"sort"
)

type User struct {
	Id int
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Id < u[j].Id
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	var users = []User{{5}, {2}, {6}, {3}, {1}, {4}}

	sort.Sort(Users(users))

	fmt.Println(users)
}