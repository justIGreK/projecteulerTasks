package main

import "fmt"

func main() {
	slice, evenSum := fibonacci(4000000)
	for _, number := range slice {
		fmt.Printf("%v ", number)
	}
	fmt.Printf("\nevenSum=%v", evenSum)
}

func fibonacci(limit int) ([]int, int) {
	fibSlice := []int{1, 2}
	evenSum := 2

	for {
		nextFib := fibSlice[len(fibSlice)-1] + fibSlice[len(fibSlice)-2]
		if nextFib > limit {
			break
		}
		fibSlice = append(fibSlice, nextFib)
		if nextFib%2 == 0 {
			evenSum += nextFib
		}
	}

	return fibSlice, evenSum
}
