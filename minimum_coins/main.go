package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumCoins(amount int, coins []int) int {
	memo := make([]int, amount+1)
	memo[0] = 0

	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, coin := range coins {
			if i >= coin {
				memo[i] = min(memo[i], memo[i-coin])
			}
		}
		memo[i]++
	}
	return memo[amount]
}

func minimumCoinsRecur(amount int, coins []int) int {
	if amount <= 0 {
		return 0
	}
	minChange := math.MaxInt32
	for _, coin := range coins {
		minChange = min(minChange, minimumCoins(amount-coin, coins))
	}
	return 1 + minChange
}

func main() {
	fmt.Println(minimumCoins(100, []int{1, 2, 3}))
	fmt.Println(minimumCoinsRecur(10, []int{1, 2, 3}))
}
