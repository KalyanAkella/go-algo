package main

import "fmt"

func solve(arr []int) int {
	for i, _ := range arr {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && l[j] > l[i] {
				l[i] = l[j]
			}
		}
		l[i] += 1
	}
	result := -1
	for _, v := range l {
		if v > result {
			result = v
		}
	}
	return result
}

var l []int

func main() {
	input := []int{50, 3, 10, 7, 40, 80}
	//input := []int{3, 2}
	//input := []int{3, 10, 2, 1, 20}
	l = make([]int, len(input))
	fmt.Println(solve(input))
}
