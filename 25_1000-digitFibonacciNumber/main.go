package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(findFibonachi(1000))
}
func findFibonachi(limit int) int {

	fibonaci1 := big.NewInt(1)
	fibonaci2 := big.NewInt(1)
	index := 2
	for {
		fibonaci1.Add(fibonaci1, fibonaci2)
		fibonaci1, fibonaci2 = fibonaci2, fibonaci1
		index++
		if len(fibonaci2.String()) >= limit {
			return index
		}
	}
}
