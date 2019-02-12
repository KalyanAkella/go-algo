package main

import "fmt"

func minLarge(arr []int, key int) int {
	ans, lo, hi := -1, 0, len(arr)-1

	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] > key {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	var result int

	result = minLarge(arr, 5)
	fmt.Println(arr[result])

	result = minLarge(arr, 70)
	fmt.Println(result)

	result = minLarge(arr, 40)
	fmt.Println(arr[result])
}
