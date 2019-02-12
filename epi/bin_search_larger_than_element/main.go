package main

import "fmt"

func search(arr []int, num int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] > num {
			hi = mid - 1
		} else if arr[mid] <= num {
			lo = mid + 1
		}
	}
	if lo < len(arr) {
		return lo
	}
	return -1
}

func main() {
	arr := []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}
	fmt.Println(search(arr, 500))
	fmt.Println(search(arr, 101))
	fmt.Println(search(arr, 285))
	fmt.Println(search(arr, 243))
	fmt.Println(search(arr, -10))
}
