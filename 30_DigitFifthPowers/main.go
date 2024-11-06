package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findSumOfDigitsNPowers(5))
}

func findSumOfDigitsNPowers(power int) int{
	maxLimit := 6 * int(math.Pow(9, float64(power))) 
	sumOfNumbers := 0

	fifthPowers := make([]int, 10)
	for i := 0; i <= 9; i++ {
		fifthPowers[i] = int(math.Pow(float64(i), float64(power)))
	}

	for i := 2; i <= maxLimit; i++ { 
		sum := 0
		number := i

	
		for number > 0 {
			digit := number % 10
			sum += fifthPowers[digit]
			number /= 10
		}


		if sum == i {
			sumOfNumbers += i
		}
	}
	return sumOfNumbers
}
