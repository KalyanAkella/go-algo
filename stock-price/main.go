package main

import "fmt"

func best_profit(prices []int) (int, int) {
	if len(prices) < 2 {
		panic("Must have at least two prices for profit computation")
	}

	max_profit := prices[1] - prices[0]
	min_price := prices[0]
	for i := 1; i < len(prices); i++ {
		current_price := prices[i]
		current_profit := current_price - min_price
		min_price = min(min_price, current_price)
		max_profit = max(max_profit, current_profit)
	}
	return max_profit, min_price
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(best_profit([]int{10, 7, 5, 8, 11, 9}))
	fmt.Println(best_profit([]int{10, 9, 8, 7, 6, 5}))
}
