package main

import (
	"fmt"
	"learn-go/05-packages-c/3-constant/date"
)

func main() {
	fmt.Println(date.DaysInWeek)
	fmt.Println(date.WeeksToDays(40))
	fmt.Println(date.DaysToWeeks(280))
}