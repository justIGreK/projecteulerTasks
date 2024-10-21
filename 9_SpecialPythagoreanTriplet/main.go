package main

import (
	"fmt"
	"math"
)

func main() {
	a,b,c :=findTripletWithSum(1000)
	fmt.Println(a*b*c)
}
func findTripletWithSum(sum float64) (float64, float64, float64) {
	for b := float64(1); ; b++ {
		for a := float64(1); a < b; a++ {
			c := math.Sqrt(a*a + b*b)
			if a+b+c == sum {
				if isPythagoreanTriplet(a, b, c) {
					return a, b, c
				}
			}
		}
	}
}

func isPythagoreanTriplet(a, b, c float64) bool {
	return a*a+b*b == c*c
}
