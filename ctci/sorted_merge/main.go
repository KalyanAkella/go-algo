package main

import "fmt"

func sortedMerge(a, b []int) []int {
	temp := make([]int, len(a)+len(b))
	var short, long []int
	if len(a) > len(b) {
		short = b
		long = a
	} else {
		short = a
		long = b
	}
	for i := range long {
		temp[i] = long[i]
	}
	merge(temp, short, len(long))
	return temp
}

func merge(long, short []int, startIndex int) {
	for i, j := len(short)-1, startIndex; i >= 0 && j < len(long); i, j = i-1, j+1 {
		long[j] = short[i]
		for k := j - 1; k >= 0; k-- {
			if long[k] > long[k+1] {
				long[k], long[k+1] = long[k+1], long[k]
			}
		}
	}
}

func main() {
	a := []int{1, 3, 5, 7, 9, 11, 13}
	b := []int{2, 4, 6, 8}
	ans := sortedMerge(a, b)
	fmt.Println(ans)
}
