package main

import (
	"fmt"
	"learn-go/ch10/packages/calendar"
	"log"
)

// type Date struct {
// 	Year  int
// 	Month int
// 	Day   int
// }

func main() {
	// date := Date{Year: 2023, Month: 10, Day: 21}
	// date = Date{Year: 0, Month: 0, Day: 0}

	date := calendar.Date{}

	fmt.Println(date.Year())

	err := date.SetYear(1989)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Year())

	fmt.Println()

	fmt.Println(date.Month())

	err = date.SetMonth(4)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Month())

	fmt.Println()

	fmt.Println(date.Day())

	err = date.SetDay(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Day())
}