package custom

import "fmt"

const PublicConstant = 1
const privateConstant = 2

var PublicVariable = 3
var privateVariable = 4

func PublicFunction() {
	const Constant = 5

	fmt.Println("Public Function:", Constant)
}

func privateFunction() {
	fmt.Println("Private Function")
}

type PublicType int
type privateType int

type PublicStruct struct {
	PublicField  int
	privateField int
}

type privateStruct struct {
	PublicField  int
	privateField int
}


