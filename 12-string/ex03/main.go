package main

import "fmt"

func main() {
	var str = "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", str[i], str[i], str[i])
	}

	fmt.Println()

	var arr = []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", arr[i], arr[i], arr[i])
	}

	fmt.Println()

	for _, value := range str {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", value, value, value)
	}
}