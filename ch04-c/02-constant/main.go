package main

import (
	"fmt"
	"learn-go/ch04-c/packages/date"
)

func main() {
	fmt.Println(date.DaysInWeek)
	fmt.Println(date.WeeksToDays(40))
	fmt.Println(date.DaysToWeeks(280))
}