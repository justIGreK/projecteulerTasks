package main

import "fmt"

func main() {
	var sum int
	rangeOfNumbers := 1000
	for i := 1; i < rangeOfNumbers; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
