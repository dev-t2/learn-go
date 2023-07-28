package main

import "fmt"

func IsRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	var price = 35000

	if price > 50000 {
		if IsRichFriend() {
			fmt.Println("신발 끈이 풀렸네???")
		} else {
			fmt.Println("n 분의 1로 내자!!!")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("신발 끈이 풀렸네?")
		} else {
			fmt.Println("n 분의 1로 내자!")
		}
	} else {
		fmt.Println("오늘은 내가 살게~")
	}
}