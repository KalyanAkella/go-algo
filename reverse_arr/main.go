package main

import "fmt"

func reverse(arr []int) {
	for beg, end := 0, len(arr)-1; beg < end && beg < len(arr) && end >= 0; {
		arr[beg], arr[end] = arr[end], arr[beg]
		beg++
		end--
	}
}

func test_reverse(arr []int) {
	fmt.Println("Before:", arr)
	reverse(arr)
	fmt.Println("After:", arr)
}

func main() {
	test_reverse([]int{1, 2, 3, 4, 5, 6})
	test_reverse([]int{1, 2, 3, 4, 5})
	test_reverse([]int{1, 2})
	test_reverse([]int{1})
}
