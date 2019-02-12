package main

import "fmt"

func pairExists(arr []int, s int) bool {
	tbl := make(map[int]struct{})
	for i := range arr {
		num := arr[i]
		if _, present := tbl[num]; present {
			return true
		} else {
			tbl[s-num] = struct{}{}
		}
	}
	return false
}

func main() {
	arr := []int{1, 4, 45, 6, 10, 8}
	fmt.Println(pairExists(arr, 16))
	fmt.Println(pairExists(arr, 100))
}
