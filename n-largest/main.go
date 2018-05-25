package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func n_largest(arr []int, n int) int {
	ans := MinInt
	for ; n > 0; n-- {
		res := MinInt
		idx := -1
		for i := 0; i < len(arr); i++ {
			if res < arr[i] {
				res = arr[i]
				idx = i
			}
		}
		if idx >= 0 {
			ans = arr[idx]
			arr[idx] = MinInt
		}
	}
	return ans
}

func main() {
	fmt.Println(n_largest([]int{3, -6, 2, 8, 7, -1}, 10))
}
