package main

import (
	"fmt"
	"math"
)

func main() {
	num := 600851475143
	result := findLargestPrime(num)
	fmt.Printf("The largest prime factor of %d is %d\n", num, result)
}

func findLargestPrime(number int) int {
	var largest int
	for number%2 == 0 {
		largest = 2
		number = number / 2
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			if isPrime(int(i)) {
				largest = i
			}
		}
	}
	if largest < 2 {
		largest = number
	}
	return largest
}
func isPrime(number int) bool {
	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}
