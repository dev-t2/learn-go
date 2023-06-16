package main

import "fmt"

func main() {
	var num1 = 10
	var num2 = 20
	var num3 = 123456789.1234

	fmt.Print("mum1: ", num1, " num2: ", num2, " num3: ", num3, "\n")

	fmt.Println("mum1:", num1, "num2:", num2, "num3:", num3)

	fmt.Printf("num1: %d num2: %d num3: %f\n", num1, num2, num3)
	fmt.Printf("num1: %d num2: %d num3: %g\n", num1, num2, num3)
}