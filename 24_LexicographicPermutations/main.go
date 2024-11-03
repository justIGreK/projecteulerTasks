package main

import "fmt"

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 1000000
	fmt.Println(getLexicographicPermutation(digits, n))
}

func getLexicographicPermutation(digits []int, n int) []int {
	permutation := []int{}
	k := n - 1 
	for len(digits) > 0 {
		fact := factorial(len(digits) - 1)
		index := k / fact
		permutation = append(permutation, digits[index])
		digits = append(digits[:index], digits[index+1:]...) 
		k %= fact
	}
	return permutation
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}


