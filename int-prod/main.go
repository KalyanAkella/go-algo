package main

import "fmt"

func int_prod(arr []int) []int {
	result := make([]int, len(arr))
	product := 1
	for i := range arr {
		result[i] = product
		product *= arr[i]
	}
	product = 1
	for i := len(arr) - 1; i >= 0; i-- {
		result[i] *= product
		product *= arr[i]
	}
	return result
}

func main() {
	fmt.Println(int_prod([]int{2, 7, 3, 4}))
	fmt.Println(int_prod([]int{1, 2, 6, 5, 9}))
	fmt.Println(int_prod([]int{1, 0, 6, 5, 9}))
}
