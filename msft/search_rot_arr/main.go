package main

import "fmt"

func isBoundary(arr []int, in int) bool {
	if in < 0 || in > len(arr)-1 {
		return false
	}
	return (in > 0 && arr[in-1] > arr[in]) || (in < len(arr)-1 && arr[in+1] < arr[in])
}

func search(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		bound := isBoundary(arr, mid)
		if arr[mid] < k {
			if bound {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if arr[mid] > k {
			if bound {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{5, 6, 7, 8, 9, 10, 1, 2, 3}, 10))
	fmt.Println(search([]int{3, 1, 2}, 1))
	fmt.Println(search([]int{3, 5, 1, 2}, 6))
}
