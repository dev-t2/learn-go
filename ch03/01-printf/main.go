package main

import "fmt"

func main() {
	fmt.Println(1.0 / 3.0)
	
	fmt.Printf("%0.2f\n", 1.0 / 3.0)

	result := fmt.Sprintf("%0.2f\n", 1.0 / 3.0)

	fmt.Printf(result)
}