package date

const DaysInWeek = 7

func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days int) int {
	return (days / DaysInWeek) + 1
}