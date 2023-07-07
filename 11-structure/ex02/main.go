package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type Owner struct {
	UserInfo User
	Level    int
}

type Customer struct {
	User
	Name  string
	Level int
}

func main() {
	var user = User{0, "Go"}
	
	fmt.Printf("아이디: %d, 이름: %s\n", user.Id, user.Name)
	
	var owner = Owner{user, 1}

	fmt.Printf("아이디: %d, 이름: %s, 레벨: %d\n", owner.UserInfo.Id, owner.UserInfo.Name, owner.Level)

	var customer = Customer{user, "Go Language", 2}
	
	fmt.Printf("아이디: %d, 닉네임: %s, 이름: %s, 레벨: %d\n", customer.Id, customer.Name, customer.User.Name, customer.Level)
}