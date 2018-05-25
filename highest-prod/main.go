package main

import "fmt"

// [7, 8, 10, 2, 1, 9, 3] ->
// can there be duplicates ?
// brute force way performs the same partial multiplication many times

func high_prod(arr []int) (int, []int) {
	if len(arr) < 3 {
		panic("Array size must at least be 3")
	}

	high_prod_3 = arr[0] * arr[1] * arr[2]
	high_prod_2 = arr[0] * arr[1]
	low_prod_2 = arr[0] * arr[1]
	high = max(arr[0], arr[1])

}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	} else {
		return c
	}
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}

func main() {
	fmt.Println(high_prod([]int{7, 8, 19, 2, 1, 9, 3}))
	fmt.Println(high_prod([]int{-10, -10, 1, 3, 2}))
	fmt.Println(high_prod([]int{1, 10, -5, 1, -100}))
}
