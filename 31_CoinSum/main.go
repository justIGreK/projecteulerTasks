package main

import "fmt"

func main() {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	amount := 200

	fmt.Println(countWaysToMake(amount, coins))
}

func countWaysToMake(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}
