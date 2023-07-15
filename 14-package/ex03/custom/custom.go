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

func (p PublicStruct) PublicMethod() {
	fmt.Println("Public Struct - Public Method")
}

func (p PublicStruct) privateMethod() {
	fmt.Println("Public Struct - Private Method")
}

type privateStruct struct {
	PublicField  int
	privateField int
}

func (p privateStruct) privateMethod() {
	fmt.Println("Private Struct - Private Method")
}
