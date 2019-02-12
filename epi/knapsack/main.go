package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValueDP(costs, weights []int, weightLimit int) int {
	numItems := len(costs)
	memo := make([][]int, numItems+1)
	for i := range memo {
		memo[i] = make([]int, weightLimit+1)
	}

	for i := 1; i <= numItems; i++ {
		currWeight := weights[i-1]
		currCost := costs[i-1]
		for j := 0; j <= weightLimit; j++ {
			if currWeight <= j {
				memo[i][j] = max(currCost+memo[i-1][j-currWeight], memo[i-1][j])
			} else {
				memo[i][j] = memo[i-1][j]
			}
		}
	}
	return memo[numItems][weightLimit]
}

func maxValue(costs, weights []int, numItems, remainingWeight int) int {
	if remainingWeight <= 0 || numItems <= 0 {
		return 0
	}

	maxResult := maxValue(costs, weights, numItems-1, remainingWeight)
	lastItemWeight := weights[numItems-1]
	if lastItemWeight <= remainingWeight {
		maxValueWithLastItem := costs[numItems-1] + maxValue(costs, weights, numItems-1, remainingWeight-lastItemWeight)
		maxResult = max(maxResult, maxValueWithLastItem)
	}
	return maxResult
}

func main() {
	costs := []int{65, 35, 245, 195, 65, 150, 275, 155, 120, 320, 75, 40, 200, 100, 220, 99}
	weights := []int{20, 8, 60, 55, 40, 70, 85, 25, 30, 65, 75, 10, 95, 50, 40, 10}
	fmt.Println(maxValue(costs, weights, 16, 130))
	fmt.Println(maxValueDP(costs, weights, 130))
}
