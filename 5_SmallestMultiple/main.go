package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSmallestMultiple(20))
}

func findSmallestMultiple(limit int) int {
	for i := 40; ; i += 2 {
		if isEvenlyDivisible(i, limit) {
			return i
		}
	}
}
func isEvenlyDivisible(number, limit int) bool {
	for i := 1; i <= limit; i++ {
		if number%i != 0 {
			return false
		}
	}
	return true
}
