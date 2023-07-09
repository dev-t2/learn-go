package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func AverageScore(math int, english int, history int) {
	var total = math + english + history
	var average = total / 3

	fmt.Println("평균 점수는", average, "점 입니다.")
}

func PrintString(nums ...int) {
	fmt.Println(nums)
}

func main() {
	var num = Add(1, 2)

	fmt.Println(num)

	fmt.Println()

	AverageScore(80, 74, 95)
	AverageScore(88, 92, 53)
	AverageScore(78, 73, 78)

	fmt.Println()

	PrintString(1, 2, 3, 4, 5)
}