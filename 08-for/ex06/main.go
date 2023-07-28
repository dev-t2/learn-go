package main

import "fmt"

func main() {
	var i, j = 1, 1

	// var isExit = false

	// for ; i <= 9; i++ {
	// 	for j = 1; j <= 9; j++ {
	// 		if i*j == 45 {
	// 			isExit = true

	// 			break
	// 		}
	// 	}

	// 	if isExit {
	// 		break
	// 	}
	// }

	OuterFor: for ; i <= 9; i++ {
		for j = 1; j <= 9; j++ {
			if i * j == 45 {
				break OuterFor
			}
		}
	}

	fmt.Printf("%d * %d = %d\n", i, j, i*j)
}