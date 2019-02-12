package main

import "fmt"

func searchRotated(arr []int, num int) int {
	return binSearchRotated(arr, num, 0, len(arr)-1)
}

//TODO: To be continued...
func binSearchRotated(arr []int, num, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)>>1
	if arr[mid] == num {
		return mid
	}
	if arr[lo] < arr[mid] {
		if arr[mid] > num {
			return binSearchRotated(arr, num, lo, mid-1)
		} else {
			return binSearchRotated(arr, num, 
		}
	}
	if arr[lo] <= num && num < arr[mid] {
		return binSearchRotated(arr, num, lo, mid-1)
	}
	return binSearchRotated(arr, num, mid+1, hi)
}

func main() {
	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	//arr := []int{16, 19, 20, 25, 1, 3, 4}
	for i := range arr {
		fmt.Println(searchRotated(arr, arr[i]))
	}
}
