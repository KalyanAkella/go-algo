package main

import (
	"fmt"
	"math"
)

func largestSum(arr []int) int {
	l := len(arr)
	memo := make([]int, l)
	memo[0] = arr[0]

	maxSum := memo[0]
	for i := 1; i < l; i++ {
		memo[i] = max(memo[i-1]+arr[i], arr[i])
		maxSum = max(maxSum, memo[i])
	}
	return maxSum
}

func largestSumTopDown(arr []int, i int) int {
	if i == 0 {
		return arr[0]
	}
	return max(largestSumTopDown(arr, i-1)+arr[i], arr[i])
}

func largestSum1(arr []int) int {
	max_sum := math.MinInt32
	run_sum := 0

	for _, n := range arr {
		if run_sum+n < 0 {
			run_sum = 0
		} else {
			run_sum += n
			max_sum = max(max_sum, run_sum)
		}
	}
	return max_sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(largestSum([]int{0, 1, 3, 0, 0, 2, 9, 7, 10}))
	fmt.Println(largestSum([]int{1, 2, -3, -4, 2, 7, -2, 3}))
}
