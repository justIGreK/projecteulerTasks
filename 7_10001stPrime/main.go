package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNstPrimeNumber(10001))
}

func findNstPrimeNumber(Nst int) int{
	counter := 0
	for i:= 2; ; i++{
		if isPrime(i){
			counter++
		}
		if counter == Nst{
			return i
		}
	} 
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}
