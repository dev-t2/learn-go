package mypackage

import "fmt"

type EmbeddedType string

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Exported Method on Embedded Type")
}

func (e EmbeddedType) unexportedMethod() {
	fmt.Println("Unexported Method on Embedded Type")
}

type MyType struct {
	EmbeddedType
}