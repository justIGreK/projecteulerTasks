package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(getSumOfDigits(getFactorial(100)))

}

func getFactorial(n int64) string {
	result := big.NewInt(1)

	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}

	return result.String()
}

func getSumOfDigits(numberStr string) int {
	sum := 0
	digitStr := ""
	for i := 0; i < len(numberStr); i++ {
		if i+1 >= len(numberStr) {
			digitStr = numberStr[i:]
		} else {
			digitStr = numberStr[i : i+1]
		}
		digit, _ := strconv.Atoi(digitStr)
		sum += digit
	}
	return sum
}
