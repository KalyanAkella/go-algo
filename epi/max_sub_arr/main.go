package main

import (
	"fmt"
	"math"
)

func maxSubArr(arr []int) (int, int, int) {
	beg, end := 0, 0
	min_sum, sum, max_sum := math.MaxInt32, 0, math.MinInt32
	min_idx := -1

	for i := range arr {
		sum += arr[i]
		if sum < min_sum {
			min_sum = sum
			min_idx = i
		}
		subSum := sum - min_sum
		if subSum > max_sum {
			max_sum = subSum
			beg, end = min_idx+1, i+1
		}
	}
	return beg, end, max_sum
}

func main() {
	arr := []int{904, 40, 523, 12, -335, -385, 124, 481, -31}
	beg, end, sum := maxSubArr(arr)
	fmt.Println(arr[beg:end], sum)
}
