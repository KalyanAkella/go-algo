package main

import (
	"fmt"
	"math"
)

/*
Questions:
1. How big is the array ?
2. How big are each of these numbers ?
3. Can there be negative numbers ?
*/
func maxSubArray(arr []int) (int, int) {
	minIndex, resultB, resultE := -1, -1, -1
	minSum, maxSum := math.MaxInt32, math.MinInt32
	sum := 0

	for i, n := range arr {
		sum += n
		if sum < minSum {
			minSum = sum
			minIndex = i
		}
		rangeSum := sum - minSum
		if rangeSum > maxSum {
			maxSum = rangeSum
			resultB, resultE = minIndex, i
		}
	}
	return resultB, resultE
}

func main() {
	arr := []int{904, 40, 523, 12, -335, -385, -124, 481, -31}
	fmt.Println(maxSubArray(arr))
}
