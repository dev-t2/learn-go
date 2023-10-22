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

	err := date.SetYear(2023)

	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(10)

	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(22)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}