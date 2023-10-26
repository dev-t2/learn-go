package main

import (
	"fmt"
	"learn-go/ch10/packages/calendar"
	"learn-go/ch10/packages/mypackage"
	"log"
)

func main() {
	value := mypackage.MyType{}

	value.ExportedMethod()

	fmt.Println()

	event := calendar.Event{}

	err := event.SetTitle("Birthday")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())

	err = event.SetYear(1989)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Year())

	err = event.SetMonth(4)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Month())

	err = event.SetDay(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Day())
}
