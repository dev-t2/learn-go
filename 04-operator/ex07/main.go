package main

import (
	"fmt"
	"math/big"
)

func main() {
	var num1, _ = new(big.Float).SetString("0.1")
	var num2, _ = new(big.Float).SetString("0.2")
	var num3, _ = new(big.Float).SetString("0.3")
	var num4 = new(big.Float).Add(num1, num2)

	fmt.Println(num1, num2, num3, num4)
	fmt.Println(num3.Cmp(num4))
}