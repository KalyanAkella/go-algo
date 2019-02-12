package main

import "fmt"

func magicIndexLinear(arr []int) int {
	for i := range arr {
		if i == arr[i] {
			return i
		}
	}
	return -1
}

func magicIndexDistinct(arr []int, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)>>1
	if arr[mid] == mid {
		return mid
	}
	if arr[mid] < mid {
		return magicIndexDistinct(arr, mid+1, hi)
	}
	return magicIndexDistinct(arr, lo, mid-1)
}

func magicIndex(arr []int, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)>>1
	if arr[mid] == mid {
		return mid
	}
	if arr[mid] < mid {
		index := max(mid+1, arr[mid])
		return magicIndexDistinct(arr, index, hi)
	}
	index := min(mid-1, arr[mid])
	return magicIndexDistinct(arr, lo, index)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	distinctArr := []int{-40, -20, -1, 1, 2, 3, 5, 7, 9, 12, 13}
	dupsArr := []int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}

	fmt.Println(magicIndexDistinct(distinctArr, 0, len(distinctArr)-1))
	fmt.Println(magicIndexDistinct(dupsArr, 0, len(dupsArr)-1))

	fmt.Println(magicIndex(distinctArr, 0, len(distinctArr)-1))
	fmt.Println(magicIndex(dupsArr, 0, len(dupsArr)-1))
}
