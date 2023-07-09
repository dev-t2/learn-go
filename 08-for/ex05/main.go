package main

import "fmt"

func main() {
	var dan = 2
	var num = 1

	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, num, dan * num)

			num++

			if num == 10 {
				break
			}
		}

		fmt.Println()

		dan++
		num = 1

		if dan == 10 {
			break
		}
	}

	fmt.Println("프로그램 종료")
}