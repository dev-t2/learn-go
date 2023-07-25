package main

import "fmt"

func main() {
	var temp [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for index, value := range temp {
		fmt.Println(index, value)
	}
}