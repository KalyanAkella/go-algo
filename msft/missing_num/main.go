package main

import "fmt"

func findMissingNum(arr []int) int {
	la := len(arr)
	tbl := make([]int, la+1)

	for i := range arr {
		tbl[arr[i]-1]++
	}

	for i := range tbl {
		if tbl[i] == 0 {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findMissingNum([]int{1, 2, 3, 5}))
	fmt.Println(findMissingNum([]int{1, 2, 3, 4, 5, 6, 7, 8, 10}))
	fmt.Println(findMissingNum([]int{1, 2, 3, 4, 5, 7, 8, 9, 10}))
	fmt.Println(findMissingNum([]int{10, 2, 3, 4, 5, 7, 8, 1, 9}))
}
