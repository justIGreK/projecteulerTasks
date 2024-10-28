package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findHighlyDivisible(500))
}


func findHighlyDivisible(count int) int {
	triangleNumber := 0
	for i := 1; ; i++ { 
		triangleNumber += i
		if tempCount := getDivisorsCount(triangleNumber); tempCount >= count {
			return triangleNumber
		}
	}
}

func getDivisorsCount(number int) int {
	count := 0
	limit := int(math.Sqrt(float64(number)))
	for i := 1; i < limit; i++ {
		if number%i == 0 {
			count += 2 
		}
	}
	return count
}
