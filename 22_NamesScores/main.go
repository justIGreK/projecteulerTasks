package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func nameValue(name string) int {
	value := 0
	for _, char := range name {
		value += int(char - 'A' + 1) 
	}
	return value
}

func main() {
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("open error:", err)
		return
	}
	defer file.Close()
	fmt.Println(findNamesScore(file))
}

func findNamesScore(file *os.File) int{
	reader := csv.NewReader(file)
	names, err := reader.Read()
	if err != nil {
		fmt.Println("reading error:", err)
		return 0
	}

	for i := range names {
		names[i] = strings.Trim(names[i], "\"")
	}
	sort.Strings(names)
	totalScore := 0
	for index, name := range names {
		score := nameValue(name) * (index + 1) 
		totalScore += score
	}
	return totalScore
}
