package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println('1')
	fmt.Println('A')
	fmt.Println('b')
	fmt.Println('가')
	fmt.Println('"')
	fmt.Println('\n')

	fmt.Println()

	fmt.Println("1")
	fmt.Println("A")
	fmt.Println("b")
	fmt.Println("가")
	fmt.Println("\"")
	// fmt.Println("\n")

	fmt.Println()

	fmt.Println(true)
	fmt.Println(false)

	fmt.Println()

	fmt.Println(15)
	fmt.Println(3.141592)

	fmt.Println()

	fmt.Println(1 + 2)
	fmt.Println(3.4 - 1.2)
	fmt.Println(3 * 4)
	fmt.Println(7.5 / 5)

	fmt.Println()

	fmt.Println(1 < 5)
	fmt.Println(1 > 5)
	fmt.Println(1 <= 5)
	fmt.Println(5 >= 5)
	fmt.Println(1 + 2 == 5)
	fmt.Println(1 + 2 != 5)

	fmt.Println()

	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(15))
	fmt.Println(reflect.TypeOf(3.141592))
}