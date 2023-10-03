package main

import "fmt"

func main() {
	fmt.Println(1.0 / 3.0)
	
	fmt.Printf("%0.2f\n", 1.0 / 3.0)

	result := fmt.Sprintf("%.2f\n", 1.0 / 3.0)

	fmt.Printf(result)

	fmt.Println()

	fmt.Printf("Float: %f\n", 3.14)
	fmt.Printf("Integer: %d\n", 15)
	fmt.Printf("String: %s\n", "Hello, Go!")
	fmt.Printf("Boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent Sign: %%\n")

	fmt.Println()

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}