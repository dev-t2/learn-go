package main

import (
	"fmt"
	"learn-go/5-packages/date"
)

func main() {
	fmt.Println(date.DaysInWeek)
	fmt.Println(date.WeeksToDays(40))
	fmt.Println(date.DaysToWeeks(280))
}