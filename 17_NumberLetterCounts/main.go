package main

import (
	"fmt"
	"strings"
)

var units = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func main() {
	total := 0
	
	for i := 1; i <= 1000; i++ {
		word := numberToWords(i)
		total += countLetters(word)
	}
	
	fmt.Println(total)
}

func countLetters(s string) int {
	count := len(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "-", ""))
	return count
}

func numberToWords(n int) string {
	if n == 1000 {
		return "one thousand"
	}
	
	result := ""
	
	if n >= 100 {
		result += units[n/100] + " hundred"
		n %= 100
		if n != 0 {
			result += " and "
		}
	}
	
	if n >= 20 {
		result += tens[n/10]
		n %= 10
		if n != 0 {
			result += "-" + units[n]
		}
	} else if n > 0 {
		result += units[n]
	}
	
	return result
}


