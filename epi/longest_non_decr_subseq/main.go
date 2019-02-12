package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengLongNonDecrSubSeqRecur(arr []int, size int) int {
	result := 0
	for i := size - 1; i >= 1; i-- {
		if arr[size-1] > arr[i] {
			subLen := lengLongNonDecrSubSeqRecur(arr, i)
			result = max(result, 1+subLen)
		}
	}
	return result
}

func lengLongNonDecrSubSeqMemo(arr []int) int {
	la := len(arr)
	memo := make([]int, la+1)

	for size := 1; size <= la; size++ {
		for i := size; i > 0; i-- {
			if arr[size-1] >= arr[i-1] {
				memo[size] = max(memo[size], 1+memo[i])
			}
		}
	}

	fmt.Println(arr)
	fmt.Println(memo)
	return memo[la]
}

func upperBound(arr []int, num int) int {
	lo, hi := 0, len(arr)-1
	result := -1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] > num {
			result = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return result
}

func lengLongNonDecrSubSeq(arr []int) int {
	var tails []int
	for i := 0; i < len(arr); i++ {
		result := upperBound(tails, arr[i])
		if result == -1 {
			tails = append(tails, arr[i])
		}
	}
	return len(tails)
}

func main() {
	arr := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9}
	//fmt.Println(lengLongNonDecrSubSeqRecur(arr, len(arr)))
	//fmt.Println(lengLongNonDecrSubSeqMemo(arr))
	//fmt.Println(lengLongNonDecrSubSeqMemo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(arr)
	fmt.Println(lengLongNonDecrSubSeq(arr))
	//arr := []int{0, 8}
	//fmt.Println(upperBound(arr, 4))
	//fmt.Println(upperBound(arr, 10))
	//fmt.Println(upperBound(arr, -1))
}
