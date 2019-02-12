package main

import (
	"fmt"
	"math"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func subArrayCloseToN(arr []int, n int) (int, int, int) {
	la := len(arr)
	sums := make([]int, la+1)
	for i := 1; i <= la; i++ {
		sums[i] = sums[i-1] + arr[i-1]
	}

	beg, end := -1, -1
	minDiff := math.MaxInt32
	for i := 1; i <= la; i++ {
		for j := 0; (j + i) <= la; j++ {
			currDiff := abs(sums[j+i] - sums[j] - n)
			if currDiff < minDiff {
				minDiff = currDiff
				beg, end = j, j+i
			}
		}
	}
	return beg, end, minDiff
}

func printSubArrCloseToN(arr []int, n int) {
	start, end, minDiff := subArrayCloseToN(arr, n)
	fmt.Println(arr[start:end], minDiff)
}

func main() {
	printSubArrCloseToN([]int{904, 40, 523, 12, -335, -385, -124, 481, -31}, 100)
	printSubArrCloseToN([]int{5, 2, 1, 8, 3, 7}, 10)
	printSubArrCloseToN([]int{5, 2, 1, 10, 8, 3, 7}, 10)
	printSubArrCloseToN([]int{5, 2, 1, 10, 8, 3, 7}, 36)
	printSubArrCloseToN([]int{5}, 5)
}
