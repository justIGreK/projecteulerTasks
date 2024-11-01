package main

import "fmt"

func main() {
	fmt.Println(getSundays(1901, 2000, 2))
}

func getSundays(start, end, dayOfWeek int) int {
	daysInMonth := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	sundayCount := 0
	for start <= end {
		for month := 1; month <= 12; month++ {
			if dayOfWeek == 0 {
				sundayCount++
			}
			if month == 2 && isLeapYear(start) {
				dayOfWeek = (dayOfWeek + 29) % 7
			} else {
				dayOfWeek = (dayOfWeek + daysInMonth[month]) % 7
			}
		}
		start++
	}
	return sundayCount
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}
