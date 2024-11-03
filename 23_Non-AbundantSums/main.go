package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findAllAbudant(28123))
}

func isAbudant(number int) bool {
	divisors := []int{1}
	limit := int(math.Sqrt(float64(number)))
	for i := 2; i <= limit; i++ {
		if number%i == 0 {
			if number/i == i{
				divisors = append(divisors, i)
			}else{
				divisors = append(divisors, i, number/i)
			}
		}
	}
	sum := 0
	for _, divisor := range divisors {
		sum += divisor
	}
	if sum > number {
		return true
	}
	return false
}

func findAllAbudant(limit int) int {
	abudants := []int{}
	nonAbudants := []int{}
	for i := 1; i <= limit; i++ {
		if isAbudant(i) {
			abudants = append(abudants, i)
		}
		if !isPosibleToWriteW2Abs(abudants, i){
			nonAbudants = append(nonAbudants, i)
		}
	}
	sum := 0
	for _, nonAbudant := range nonAbudants {
		sum += nonAbudant
	}
	return sum
}

func isPosibleToWriteW2Abs(abudants []int, check int)bool{
	length := len(abudants)-1
	for i:= 0; i < length; i++{
		for j:=0; j<length; j++{
			if abudants[i] + abudants[j] == check{
				return true 
			}
		}
	}
	return false
}