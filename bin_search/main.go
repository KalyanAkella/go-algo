package main

import "fmt"

func bin_search(arr []int, find, lo, hi int) int {
	fmt.Println(lo, hi)
	if lo >= hi {
		if lo < len(arr) && arr[lo] == find {
			return lo
		} else {
			return -1
		}
	}
	mid := (lo + hi) >> 1
	if arr[mid] == find {
		return mid
	} else if arr[mid] < find {
		return bin_search(arr, find, mid+1, hi)
	} else {
		return bin_search(arr, find, lo, mid-1)
	}
}

func main() {
	arr := []int{5, 8, 11, 16, 17}
	fmt.Println(16, bin_search(arr, 16, 0, len(arr)))
	fmt.Println(17, bin_search(arr, 17, 0, len(arr)))
	fmt.Println(8, bin_search(arr, 8, 0, len(arr)))
	fmt.Println(4, bin_search(arr, 4, 0, len(arr)))
	fmt.Println(18, bin_search(arr, 18, 0, len(arr)))
	fmt.Println(18, bin_search([]int{}, 18, 0, len([]int{})))
}
