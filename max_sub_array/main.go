package main

import (
	"fmt"
	"math"
)

func max_sub_arr_slow(arr []int) []int {
	max_sum := math.MinInt32
	var result []int
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			if sum > max_sum {
				max_sum = sum
				result = arr[i:(j + 1)]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(max_sub_arr_slow([]int{1, -3, 2, 1, -1}))
	fmt.Println(max_sub_arr_slow([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
