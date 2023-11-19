package main

import "fmt"

func main() {
	fmt.Println(1.0 / 3.0)
	
	fmt.Printf("%0.2f\n", 1.0 / 3.0)

	result := fmt.Sprintf("%.2f\n", 1.0 / 3.0)

	fmt.Printf(result)

	fmt.Println()

	fmt.Printf("int: %d\n", 15)
	fmt.Printf("float: %f\n", 3.141592)
	fmt.Printf("string: %s\n", "Hello, Go!")
	fmt.Printf("bool: %t\n", false)
	fmt.Printf("values: %v %v %v %v\n", 1.2, 'a', "\t", true)
	fmt.Printf("values: %#v %#v %#v %#v\n", 1.2, 'a', "\t", true)
	fmt.Printf("types: %T %T %T %T\n", 1.2, 'a', "\t", true)

	fmt.Println()

	fmt.Printf("%6.3f\n", 12.3456)
	fmt.Printf("%6.2f\n", 12.3456)
	fmt.Printf("%6.1f\n", 12.3456)
	fmt.Printf("%.1f\n", 12.3456)
	fmt.Printf("%.2f\n", 12.3456)
}