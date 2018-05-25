package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func min_max(arr []int) (int, int) {
	min := MaxInt
	max := MinInt
	for _, n := range arr {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func sort(arr []int, min int, max int) []int {
	res := make([]int, len(arr))
}

func main() {
	arr := []int{3, 7, 4, 8, 6, 2}
	fmt.Println(min_max(arr))
}
