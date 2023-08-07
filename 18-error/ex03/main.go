package main

import (
	"fmt"
)

type PasswordError struct {
	Len int
}

func (error PasswordError) Error() string {
	return "Password Error"
}

func CreateUser(email, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password)}
	}

	return nil
}

func main() {
	var error = CreateUser("Email", "Password")

	if error != nil {
		if passwordError, isSuccess := error.(PasswordError); isSuccess {
			fmt.Printf("%v, Len: %d\n", passwordError, passwordError.Len)
		}
	} else {
		fmt.Println("Create User Success")
	}
}