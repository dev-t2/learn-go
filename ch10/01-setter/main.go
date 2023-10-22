package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	date := Date{Year: 2023, Month: 10, Day: 21}

	fmt.Println(date)

	date = Date{Year: 0, Month: 0, Day: 0}

	fmt.Println(date)

	fmt.Println()

	date = Date{}

	date.SetYear(0)
	date.SetMonth(0)
	date.SetDay(0)

	fmt.Println(date)
}