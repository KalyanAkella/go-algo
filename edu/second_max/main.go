package main

import (
	"fmt"
	"math"
)

func findSecondMax(arr []int) int {
	first_max, second_max := arr[0], math.MinInt32
	for i := 1; i < len(arr); i++ {
		if arr[i] > first_max {
			first_max, second_max = arr[i], first_max
		} else if arr[i] < first_max && arr[i] > second_max {
			second_max = arr[i]
		}
	}
	return second_max
}

func main() {
	fmt.Println(findSecondMax([]int{355, 872, 412, 872, 412, 11, 355, 10}))
	fmt.Println(findSecondMax([]int{9, 8, 7, 6, 5, 4}))
	fmt.Println(findSecondMax([]int{4, 5, 6, 7, 8, 9}))
	fmt.Println(findSecondMax([]int{12, 35, 1, 10, 34, 1}))
	fmt.Println(findSecondMax([]int{10, 5, 10}))
	fmt.Println(findSecondMax([]int{10, 10, 10}))
}
