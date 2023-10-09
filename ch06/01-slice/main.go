package main

import "fmt"

func main() {
	// var notes []string
	// notes = make([]string, 7)

	notes := make([]string, 7)

	notes[0] = "do"
	notes[3] = "fa"
	notes[6] = "ti"

	fmt.Println(notes[0], notes[3], notes[6])

	fmt.Println()

	notes = []string{"do", "re", "mi", "fa", "so", "la", "ti"}

	fmt.Println(notes[0], notes[3], notes[6])

	fmt.Println()

	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	fmt.Println()

	for _, note := range notes {
		fmt.Println(note)
	}

	fmt.Println()

	underlyingArray := [5]string{"a", "b", "c", "d", "e"}

	slice1 := underlyingArray[0:3]

	fmt.Println(slice1)

	slice2 := underlyingArray[2:5]

	fmt.Println(slice2)

	slice3 := underlyingArray[:3]

	fmt.Println(slice3)

	slice4 := underlyingArray[1:]

	fmt.Println(slice4)

	fmt.Println()

	underlyingArray[1] = "x"

	fmt.Println(underlyingArray)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Println()

	slice1[2] = "y" 

	fmt.Println(underlyingArray)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}