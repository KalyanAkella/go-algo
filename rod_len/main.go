package main

import "fmt"

func maxPriceMemo(prices, memo []int, length int) int {
	if memo[length] == 0 {
		ans := 0

		for i := 0; i < length; i++ {
			ans = max(ans, prices[i]+maxPriceMemo(prices, memo, length-i-1))
		}
		memo[length] = ans
	}
	return memo[length]
}

func maxPrice(prices []int, length int) int {
	ans := 0

	for i := 0; i < length; i++ {
		ans = max(ans, prices[i]+maxPrice(prices, length-i-1))
	}
	return ans
}

func maxPriceOld(prices []int) int {
	rodLength := len(prices)
	memo := make([]int, rodLength)
	memo[0] = prices[0]

	for l := 1; l < rodLength; l++ {
		for x := 0; x < l; x++ {
			memo[l] = max(prices[l], memo[x]+memo[l-x])
		}
	}
	fmt.Println(memo)
	return memo[rodLength-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	/*
		rodLength := 1000
		for i := 0; i < rodLength; i++ {
		}*/
	prices := []int{3, 7, 8, 9, 12, 13, 19, 20}
	fmt.Println(maxPrice(prices, len(prices)))
	fmt.Println(maxPriceMemo(prices, make([]int, len(prices)+1), len(prices)))
}
