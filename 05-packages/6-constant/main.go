package main

import (
	"fmt"
	date "learn-go/05-packages/5-date"
)

func main() {
	fmt.Println(date.DaysInWeek)
	fmt.Println(date.WeeksToDays(40))
	fmt.Println(date.DaysToWeeks(280))
}