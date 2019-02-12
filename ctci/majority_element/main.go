package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func hasMajority(arr []int, num int) bool {
	count := 0
	for _, n := range arr {
		if num == n {
			count++
		}
	}
	return count > len(arr)/2
}

func majorityElementSlow(arr []int) int {
	for _, n := range arr {
		if hasMajority(arr, n) {
			return n
		}
	}
	return -1
}

func getCandidate(arr []int) int {
	count := 0
	var candidate int
	for _, n := range arr {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElement(arr []int) int {
	candidate := getCandidate(arr)
	if hasMajority(arr, candidate) {
		return candidate
	}
	return -1
}

func majorityElementWrong(arr []int, lo, hi int) int {
	fmt.Println(lo, hi)
	larr := hi - lo + 1
	if larr < 2 {
		return -1
	} else if larr == 2 {
		in := min(lo, hi)
		if arr[in] == arr[in+1] {
			return arr[in]
		} else {
			return -1
		}
	} else {
		mid := lo + (hi-lo)>>1
		majSub1 := majorityElementWrong(arr, lo, mid)
		majSub2 := majorityElementWrong(arr, mid+1, hi)
		if (majSub1 == -1 && majSub2 == -1) || majSub1 != majSub2 {
			return -1
		} else if min(majSub1, majSub2) == -1 {
			return max(majSub1, majSub2)
		} else if majSub1 == majSub2 {
			return majSub1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	fmt.Println(majorityElement(arr))
}
