package main

import "fmt"

func prodElements(input []int) []int {
	result := make([]int, len(input)+1)
	result[0] = 1
	for i := 1; i <= len(input); i++ {
		result[i] = result[i-1] * input[i-1]
	}
	return result
}

func main() {
	fmt.Println(prodElements([]int{1, 2, 3, 4, 5}))
}
