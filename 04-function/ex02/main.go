package main

import "fmt"

func AverageScore(math int, english int, history int) {
	var total = math + english + history
	var average = total / 3

	fmt.Println("평균 점수는", average, "점 입니다.")
}

func main() {
	AverageScore(80, 74, 95)
	AverageScore(88, 92, 53)
	AverageScore(78, 73, 78)
}