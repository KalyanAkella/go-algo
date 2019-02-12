package main

import "fmt"

/*
	[4 5 6 7 0 1 2], 4 -> 0
	[1 2 4 5 6 7 0]
*/

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	for _, n := range input {
		fmt.Println(searchRotatedSortedArr(input, n))
	}
}

func findPivot(arr []int) int {
	s, e := 0, len(arr)-1
	for s <= e {
		mid := s + (e-s)>>1
		left := mid - 1  // TODO: Bounds check
		right := mid + 1 // TODO: Bounds check
		if arr[left] < arr[mid] && arr[mid] > arr[right] {
			return mid
		} else if arr[left] < arr[mid] && arr[mid] < arr[right] {
			s = right
		} else if arr[left] > arr[mid] && arr[mid] > arr[right] {
			e = left
		}
	}
	return -1
}

func find(arr []int, target, start, end int) int {
	for start <= end {
		mid := start + (end-start)>>1
		if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func searchRotatedSortedArr(arr []int, target int) int {
	pivotIndex := findPivot(arr)
	if pivotIndex == -1 {
		return -1
	}
	ans := find(arr, target, 0, pivotIndex)
	if ans == -1 {
		return find(arr, target, pivotIndex+1, len(arr)-1)
	}
	return ans
}
