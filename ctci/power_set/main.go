package main

import "fmt"

func powerSet(arr []int, size int) [][]int {
	var ans [][]int
	if size == 0 {
		ans = make([][]int, 1)
		ans[0] = make([]int, 0)
		return ans
	}
	sub := powerSet(arr, size-1)
	curr := arr[size-1]
	ans = append(ans, sub...)
	for i := range sub {
		sub[i] = append(sub[i], curr)
	}
	ans = append(ans, sub...)
	return ans
}

func main() {
	fmt.Println(powerSet([]int{1, 2, 3, 4, 5}, 5))
}
