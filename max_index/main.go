package main

import "fmt"

/*
Given an array A of positive integers. The task is to find the maximum of j - i subjected to the constraint of A[i] <= A[j].

Eg:
arr = [34 8 10 3 2 80 30 33 1]

Return: 6
Expl: Since A[1] < A[7] and 7-1 is 6
*/

func maxIndexSlo(arr []int) int {
	ii, jj := 0, 0
	maxDiff := 0
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] <= arr[j] && (j-i) >= maxDiff {
				maxDiff = j - i
				ii, jj = i, j
			}
		}
	}
	return jj - ii
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxIndex(arr []int) int {
	la := len(arr)
	lmin := make([]int, la)
	rmax := make([]int, la)

	lmin[0] = arr[0]
	for i := 1; i < la; i++ {
		lmin[i] = min(arr[i], lmin[i-1])
	}

	rmax[la-1] = arr[la-1]
	for i := la - 2; i >= 0; i-- {
		rmax[i] = max(arr[i], rmax[i+1])
	}

	maxDiff := 0
	for i, j := 0, 0; i < la && j < la; {
		if lmin[i] < rmax[j] {
			maxDiff = max(j-i, maxDiff)
			j++
		} else {
			i++
		}
	}
	return maxDiff
}

func main() {
	arr := []int{34, 8, 10, 3, 2, 80, 30, 33, 1}
	fmt.Println(maxIndex(arr))
}
