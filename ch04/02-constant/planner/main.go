package main

import (
	"fmt"
	"learn-go/ch04/02-constant/dates"
)

func main() {
	days := 3

	fmt.Println(days)
	fmt.Println(days + dates.DaysInWeek)
}