package main

import (
	"fmt"
	"learn-go/ch10/packages/calendar"
	"log"
)



func main() {
	date := calendar.Date{Year: 2023, Month: 10, Day: 21}

	fmt.Println(date)

	fmt.Println()

	date = calendar.Date{Year: 0, Month: 0, Day: 0}

	fmt.Println(date)

	fmt.Println()

	date = calendar.Date{}

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