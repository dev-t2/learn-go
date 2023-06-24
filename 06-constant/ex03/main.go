package main

import "fmt"

// const Pig = 0
// const Cow = 1
// const Chicken = 2

// const (
// 	 Pig = iota
//   Cow = iota
//   Chicken = iota
// )

const (
	Pig = iota
  Cow
  Chicken
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}