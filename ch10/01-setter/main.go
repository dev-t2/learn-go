package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{Year: 2023, Month: 10, Day: 21}

	fmt.Println(date)

	date = Date{Year: 0, Month: 0, Day: 0}

	fmt.Println(date)
}