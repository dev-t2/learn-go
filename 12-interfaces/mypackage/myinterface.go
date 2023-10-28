package mypackage

import "fmt"

type MyInterface interface {
	MethodWithoutParameter()
	MethodWithParameter(float64)
	MethodWithReturn() string
}

type MyInt int

func (m MyInt) MethodWithoutParameter() {
	fmt.Println("Method Without Parameters")
}

func (m MyInt) MethodWithParameter(f float64) {
	fmt.Println("Method With Parameter", f)
}

func (m MyInt) MethodWithReturn() string {
	return "Method With Return"
}

func (m MyInt) MethodNotInterface() {
	fmt.Println("Method Not interface")
}