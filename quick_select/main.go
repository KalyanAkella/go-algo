package main

import (
	"fmt"
)

/*
Returns the kth smallest element in an array from left to right
*/
func quickSelect(arr []int, left, right, k int) int {
	for left < right {
		pivotIndex := partition(arr, left, right, left+(right-left)>>1)
		if k < pivotIndex {
			right = pivotIndex - 1
		} else if k > pivotIndex {
			left = pivotIndex + 1
		} else {
			return arr[k]
		}
	}
	return arr[left]
}

func partition(arr []int, left, right, pivotIndex int) int {
	pivot := arr[pivotIndex]
	arr[right], arr[pivotIndex] = arr[pivotIndex], arr[right]
	storeIndex := left
	for i := left; i <= right-1; i++ {
		if arr[i] < pivot {
			arr[storeIndex], arr[i] = arr[i], arr[storeIndex]
			storeIndex++
		}
	}
	arr[right], arr[storeIndex] = arr[storeIndex], arr[right]
	return storeIndex
}

func main() {
	arr := []int{12, 6, 8, 10, 7, 3, 5, 9, 1}
	la := len(arr)
	fmt.Println(arr)
	for i := range arr {
		fmt.Println(quickSelect(arr, 0, la-1, i))
	}
}
