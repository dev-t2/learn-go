package main

import "fmt"

func main() {
	// var notes []string
	// notes = make([]string, 7)

	notes := make([]string, 7)

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0], notes[1], notes[2])

	fmt.Println()

	fmt.Println(len(notes))

	fmt.Println()

	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	for _, note := range notes {
		fmt.Println(note)
	}

	notes = []string{
		"do", 
		"re", 
		"mi", 
		"fa", 
		"so", 
		"la", 
		"ti",
	}

	fmt.Println(notes[0], notes[3], notes[6])

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
}