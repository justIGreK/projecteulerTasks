package main

import (
	"fmt"
	"math/big"
)

func main() {
fmt.Println(findUniquePowers(2, 100))
}

func findUniquePowers(lowB, highB int) int {
	uniquePowers := make(map[string]int)
	for a := lowB; a <= highB; a++ {
		for b := lowB; b <= highB; b++ {
			base := big.NewInt(int64(a))
			result := new(big.Int).Exp(base, big.NewInt(int64(b)), nil)
			if _, found := uniquePowers[result.String()]; !found{
				uniquePowers[result.String()] = 1
			}
		}
	}
	return len(uniquePowers)
}
