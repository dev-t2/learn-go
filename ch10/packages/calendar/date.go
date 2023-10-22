package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1989 || year > 2057 {
		return errors.New("Invalid Year")
	}

	d.year = year

	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}

	d.month = month

	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}

	d.day = day

	return nil
}

func (d *Date) Day() int {
	return d.day
}