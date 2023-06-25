package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Green
	Blue
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Undefined"
	}
}

func getColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("Favorite color is", colorToString(getColor()))
}