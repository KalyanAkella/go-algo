package main

import "fmt"

var memo []int

func solve(input []int) int {
	for i, _ := range input {
		if i > 0 && input[i] > input[i-1] {
			memo[i] = memo[i-1] + 1
		}
	}

	var largest int
	for _, n := range memo {
		if n > largest {
			largest = n
		}
	}
	return largest
}

func main() {
	input := []int{3, 2, 4, 5, 6, 1, 7}
	input = []int{3}
	input = []int{1, 2, 3, 4, 0, 5, 6, 7, 8, 9, 10, 12, 1, 2, 3}
	input = []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	memo = make([]int, len(input))
	for i, _ := range memo {
		memo[i] = 1
	}
	result := solve(input)
	fmt.Println(result)
}
