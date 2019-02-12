package main

import "fmt"

func lenLongestSubArrayWithSum(arr []int, size, k int, sums []int) (int, int) {
	if size == 0 {
		return -1, -1
	}
	si, sj := lenLongestSubArrayWithSum(arr, size-1, k, sums)
	if arr[sj+1] == arr[size-1] { // cont. indices
		for i := si; i >= 0; i-- {
			subSum := sums[size-1] - sums[i]
			if subSum > k {
				return i, size - 1
			}
		}
	}
	if (sj - si + 1) > 1 {
		return si, sj
	}
	return size - 1, size - 1
}

func longestSubArrayWithSum(arr []int, k int) (int, int) {
	la := len(arr)
	sums := make([]int, la+1)
	for i := 1; i <= la; i++ {
		sums[i] = sums[i-1] + arr[i-1]
	}
	fmt.Println(sums)
	return lenLongestSubArrayWithSum(arr, la, k, sums)
}

// TODO
func main() {
	arr := []int{431, -15, 639, 342, -14, 565, -924, 635, 167, -70}
	fmt.Println(longestSubArrayWithSum(arr, 184))
}
