package main

func three() {
	panic("This call stack's too deep for me")
}

func two() {
	three()
}

func one() {
	two()
}

func main() {
	one()
}