package main

import "fmt"

func main() {
	fmt.Println(findLongestCollatzSeq(1, 1000000))
}

func findLongestCollatzSeq(start, end int) int {
	var startingNumber int
	maxLen := 0
	noChanges := 0
	for i := end - 1; i > start; i-- {
		len := getLenCollatzSeq(i)
		if len > maxLen {
			startingNumber = i
			maxLen = len
			noChanges = 0
		} else {
			noChanges++
		}
	}
	return startingNumber
}

func getLenCollatzSeq(startingNumber int) int {
	count := 1
	for startingNumber > 1 {
		if startingNumber%2 == 0 {
			startingNumber = startingNumber / 2
		} else {
			startingNumber = 3*startingNumber + 1
		}
		count++
	}
	return count
}
