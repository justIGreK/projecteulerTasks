package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSumSquareDifference(100))
}
func findSumSquareDifference(limit int) int {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= limit; i++ {
		sumOfSquares += i * i
		sum += i
	}
	squareOfSum := sum * sum
	fmt.Printf("%v - %v = ", squareOfSum, sumOfSquares)
	return squareOfSum - sumOfSquares
}
