package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var stdin = bufio.NewReader(os.Stdin)

	var num1 int
	var num2 int

	var n, err = fmt.Scanln(&num1, &num2)

	if err == nil {
		fmt.Println(n, num1, num2)
	} else {
		fmt.Println(n, err)

		stdin.ReadString('\n')
	}

	n, err = fmt.Scanln(&num1, &num2)

	if err == nil {
		fmt.Println(n, num1, num2)
	} else {
		fmt.Println(n, err)
	}
}