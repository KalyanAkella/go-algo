package main

import (
	"fmt"
	"math"
)

func biggestIndex(arr []int, a, b, c int) int {
	aVal := math.MinInt32
	bVal, cVal := aVal, aVal

	if a >= 0 && a < len(arr) {
		aVal = arr[a]
	}

	if b >= 0 && b < len(arr) {
		bVal = arr[b]
	}

	if c >= 0 && c < len(arr) {
		cVal = arr[c]
	}

	maxVal := max(aVal, max(bVal, cVal))
	if maxVal == aVal {
		return a
	}
	if maxVal == bVal {
		return b
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func peaksValleys(arr []int) {
	for i := 1; i < len(arr); i += 2 {
		j := biggestIndex(arr, i-1, i, i+1)
		if j != i {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
}

func main() {
	arr := []int{5, 8, 6, 2, 3, 4, 6}
	fmt.Println(arr)
	peaksValleys(arr)
	fmt.Println(arr)
}
