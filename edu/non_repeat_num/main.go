package main

import "fmt"

func findFirstNonRepeatNum(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	lookup := make([]int, max+1)
	for i := range arr {
		lookup[arr[i]]++
	}
	for i := range arr {
		if lookup[arr[i]] == 1 {
			return arr[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(findFirstNonRepeatNum([]int{355, 872, 412, 872, 412, 11, 355, 10}))
	fmt.Println(findFirstNonRepeatNum([]int{355, 872, 412, 872, 412, 11, 355, 11}))
}
