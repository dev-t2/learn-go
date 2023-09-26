package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println('A')
	fmt.Println('b')
	fmt.Println('가')
	fmt.Println('\t')
	fmt.Println('\n')
	fmt.Println('\\')

	fmt.Println()

	fmt.Println(true)
	fmt.Println(false)

	fmt.Println()

	fmt.Println(42)
	fmt.Println(3.141592)

	fmt.Println()

	fmt.Println(1 + 2)
	fmt.Println(5.4 - 2.2)
	fmt.Println(3 * 4)
	fmt.Println(7.5 / 5)

	fmt.Println()

	fmt.Println(4 < 6)
	fmt.Println(4 > 6)
	fmt.Println(2 + 2 == 5)
	fmt.Println(2 + 2 != 5)
	fmt.Println(4 <= 6)
	fmt.Println(4 >= 4)

	fmt.Println()

	fmt.Println(reflect.TypeOf("Hello, Go!"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.141592))
}