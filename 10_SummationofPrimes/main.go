package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findSumOfPrimesBelow(2000000))
}

func findSumOfPrimesBelow(border int) int{
	sum := 0
	if border<2{
		return sum
	}
	for i := 2; i < border; i++ {
		if isPrime(i){
			sum += i
		}
	}
	return sum
}
func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}
