package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1989 || year > 2057 {
		return errors.New("Invalid Year")
	}

	d.Year = year

	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}

	d.Month = month

	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}

	d.Day = day

	return nil
}

func main() {
	date := Date{Year: 2023, Month: 10, Day: 21}

	fmt.Println(date)

	date = Date{Year: 0, Month: 0, Day: 0}

	fmt.Println(date)

	fmt.Println()

	date = Date{}

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