package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(findDigitSum(2, 1000))
}

func findDigitSum(number, power int64) int{
	num := big.NewInt(number)
	exp := big.NewInt(power)
	result := new(big.Int).Exp(num, exp, nil)
	var digits []int
	var digitStr string
	str := result.String()
	fmt.Println(result)
	for i:=0; i<len(str); i++{
		if i+1>=len(str){
			digitStr = str[i:]
		}else{
			digitStr = str[i:i+1]
		}
		digit, _ := strconv.Atoi(digitStr)
		digits = append(digits, digit)
	} 
	sum := 0
	for _, digit := range digits{
		sum += digit
	}
	return sum
}