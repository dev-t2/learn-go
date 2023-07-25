package main

import "fmt"

func PrintValue(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("int: %d\n", int(t))
	case float64:
		fmt.Printf("float64: %f\n", float64(t))
	case string:
		fmt.Printf("string: %s\n", string(t))
	default:
		fmt.Printf("Not Supported %T: %v\n", t, t)
	}
}

type User struct {
	Id int
}

func main() {
	PrintValue(10)
	PrintValue(3.14)
	PrintValue("Hello World")
	PrintValue(User{1})
}