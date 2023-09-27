package main

import "fmt"

func main() {
	fmt.Println(1.0 / 3.0)
	
	fmt.Printf("%0.2f\n", 1.0 / 3.0)

	result := fmt.Sprintf("%0.2f\n", 1.0 / 3.0)

	fmt.Printf(result)

	fmt.Println()

	fmt.Printf("Float: %f\n", 3.14)
	fmt.Printf("Integer: %d\n", 15)
	fmt.Printf("String: %s\n", "Hello, Go!")
	fmt.Printf("Boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent Sign: %%\n")
}