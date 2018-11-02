package main

import "fmt"

/*
[5, 7, 7, 8, 8, 10], 8 -> [3, 4]
[5, 7, 7, 7, 7, 10], 7 -> [1, 4]
*/

func main() {
	fmt.Println(rangeSearch([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(rangeSearch([]int{5, 7, 7, 7, 7, 10}, 7))
	fmt.Println(rangeSearch([]int{5, 7, 7, 7, 7, 10}, 9))
}

func rangeSearch(arr []int, target int) []int {
	initialIndex := find(arr, target)
	if initialIndex == -1 {
		return []int{-1, -1}
	}
	beginIndex := findLeft(arr, target, initialIndex)
	endIndex := findRight(arr, target, initialIndex)
	return []int{beginIndex, endIndex}
}

func findRight(arr []int, target, startIndex int) int {
	s, e := startIndex, len(arr)-1
	for s <= e {
		mid := s + (e-s)>>1
		if arr[mid] < target {
			s = mid + 1
		} else if arr[mid] > target {
			e = mid - 1
		} else {
			rightIndex := findRight(arr, target, startIndex+1)
			if rightIndex == -1 {
				return startIndex
			} else {
				return rightIndex
			}
		}
	}
	return -1
}

func findLeft(arr []int, target, startIndex int) int {
	s, e := 0, startIndex
	for s <= e {
		mid := s + (e-s)>>1
		if arr[mid] < target {
			s = mid + 1
		} else if arr[mid] > target {
			e = mid - 1
		} else {
			leftIndex := findLeft(arr, target, startIndex-1)
			if leftIndex == -1 {
				return startIndex
			} else {
				return leftIndex
			}
		}
	}
	return -1
}

func find(arr []int, target int) int {
	s, e := 0, len(arr)-1
	for s <= e {
		mid := s + (e-s)>>1
		if arr[mid] < target {
			s = mid + 1
		} else if arr[mid] > target {
			e = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
