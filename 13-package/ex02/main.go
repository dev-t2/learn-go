package main

import (
	"fmt"
	"learn-go/13-package/ex02/custom"

	"github.com/guptarohit/asciigraph"
)

func main() {
	custom.CustomPrint()

	fmt.Println()

	var data = []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	var graph = asciigraph.Plot(data)

	fmt.Println(graph)
}