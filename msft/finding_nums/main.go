package main

import "fmt"

func findNums(arr []int) (int, int) {
	tbl := make(map[int]struct{})
	for i := range arr {
		num := arr[i]
		if _, present := tbl[num]; present {
			delete(tbl, num)
		} else {
			tbl[num] = struct{}{}
		}
	}
	if lt := len(tbl); lt != 2 {
		panic("Invalid input")
	}
	var result []int
	for k, _ := range tbl {
		result = append(result, k)
	}
	return min(result[0], result[1]), max(result[0], result[1])
}

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

func main() {
	fmt.Println(findNums([]int{1, 2, 3, 2, 1, 4}))
	fmt.Println(findNums([]int{2, 1, 3, 2}))
}
