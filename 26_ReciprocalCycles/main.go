package main

import "fmt"

func main() {
	fmt.Println(findLongestRecurCycle(1000))
}

func findLongestRecurCycle(limit int) int {
	maxLength := 0
	result := 0
	for i := 2; i < limit; i++ {
		length := findCycleLength(i)
		if length > maxLength {
			maxLength = length
			result = i
		}
	}
	return result
}

func findCycleLength(number int) int {
	remainders := make(map[int]int)
	value := 1
	position := 0

	for {
		if value == 0 {
			return 0
		}
		if pos, found := remainders[value]; found {
			return position - pos
		}
		remainders[value] = position
		value = (value * 10) % number
		position++
	}
}
