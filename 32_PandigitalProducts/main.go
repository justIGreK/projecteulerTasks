package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(findSumOfAllPandigitals())
}

func findSumOfAllPandigitals()int{
	sum := 0
	products := make(map[int]bool) 
	for i:= 1; i <= 100; i++{
		for j:= 1; j <= 10000; j++{
			ij := i * j
			str := fmt.Sprintf("%d%d%d", i, j, ij)
			if len(str) == 9{
				if isPandigital(str){
					if !products[ij] {
						products[ij] = true
						sum += ij
					}
				}
			}
		}
	} 
	return sum
}

func isPandigital(str string) bool{
	seen := make(map[rune]bool)
	if len(str) != 9 {
		return false
	}
	for _, ch := range str {
		if !unicode.IsDigit(ch) || ch == '0' {
			return false
		}
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}

	return true
}
