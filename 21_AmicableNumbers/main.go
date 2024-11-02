package main

import (
	"fmt"
	"math"
)

func main() {
	 fmt.Println(findSumOfAmicableNumbers(10000))

}

func findSumOfAmicableNumbers(limit int) int{
	amicables := []int{}
	for i:=2; i<=limit; i++{
		if isAmicableNumbers(i){
			amicables = append(amicables, i)
		}
	}
	sum := 0
	for _, number := range amicables{
		sum += number
	}
	fmt.Println(amicables)
	return sum
}

func isAmicableNumbers(number int) bool{
	result := findSumOfAllDivisors(number)
	if result == number{
		return false
	}
	resultOfResult := findSumOfAllDivisors(result)
	if resultOfResult == number{
		return true
	}
	return false
}

func findSumOfAllDivisors(number int) int{
	divisors := []int{1}
	limit := int(math.Sqrt(float64(number)))
	for i := 2; i <= limit; i++ {
		if number%i == 0 {
			divisors = append(divisors, i, number/i)

		}
	}
	sum := 0
	for _, divisor := range divisors{
		sum += divisor
	} 
	return sum
}
