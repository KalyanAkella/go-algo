package main

import "fmt"

func search(arr []int, num int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] > num {
			hi = mid - 1
		} else if arr[mid] < num {
			lo = mid + 1
		} else {
			if mid > 0 && arr[mid-1] == num {
				hi = mid - 1
			} else {
				return mid
			}
		}
	}
	return hi
}

func main() {
	arr := []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}
	fmt.Println(search(arr, 108))
	fmt.Println(search(arr, 285))
	fmt.Println(search(arr, 243))
}
