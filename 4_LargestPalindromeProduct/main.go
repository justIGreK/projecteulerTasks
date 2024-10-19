package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(number int) bool {
	numberStr := strconv.Itoa(number)
	reversedNumberStr := reversedNumberStr(numberStr)
	return reversedNumberStr == numberStr
}

func reversedNumberStr(str string) string {
	byteSlice := []byte(str)
	var reverseStr []byte
	for i := range byteSlice {
		reverseStr = append(reverseStr, byteSlice[len(byteSlice)-1-i])
	}
	return string(reverseStr)
}


func findLargestPalindrome() int{
	maxPalindrome := 0 
	for i:=999; i>100; i--{
		for j:=999; j>100; j--{
			check := i * j
			if check < maxPalindrome{
				break
			}
			if isPalindrome(check){
				maxPalindrome = check
			}
		}
	}
	return maxPalindrome 
}



func main() {
	fmt.Println(findLargestPalindrome())
}
