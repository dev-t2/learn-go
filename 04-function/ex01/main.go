package main

import "fmt"

func plus(num1 int, num2 int) int {
	return num1 + num2
}

func averageScore(math int, english int, history int) {
	var total = math + english + history
	var average = total / 3

	fmt.Println("평균 점수는", average, "점 입니다.")
}

func main() {
	var result = plus(1, 2)

	fmt.Println(result)

	fmt.Println()

	averageScore(80, 74, 95)
	averageScore(88, 92, 53)
}