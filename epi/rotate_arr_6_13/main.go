package main

import "fmt"

func reverse(arr []int, s, e int) {
	for ; s <= e; s, e = s+1, e-1 {
		arr[s], arr[e] = arr[e], arr[s]
	}
}

func rotateRight(arr []int, k int) {
	la := len(arr)
	if k < 0 || k > la {
		panic("invalid rotation")
	}

	reverse(arr, 0, la-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, la-1)
}

func rotateLeft(arr []int, k int) {
	la := len(arr)
	if k < 0 || k > la {
		panic("invalid rotation")
	}

	reverse(arr, 0, k-1)
	reverse(arr, k, la-1)
	reverse(arr, 0, la-1)
}

func main() {
	var arr []int
	arr = []int{3, 5, 2, 8, 6}
	fmt.Println(arr)
	rotateLeft(arr, 1)
	fmt.Println(arr)
	rotateLeft(arr, 1)
	fmt.Println(arr)
	arr = []int{3, 5, 2, 8, 6}
	rotateLeft(arr, 2)
	fmt.Println(arr)

	fmt.Println()

	arr = []int{3, 5, 2, 8, 6}
	fmt.Println(arr)
	rotateRight(arr, 1)
	fmt.Println(arr)
	rotateRight(arr, 1)
	fmt.Println(arr)
	arr = []int{3, 5, 2, 8, 6}
	rotateRight(arr, 2)
	fmt.Println(arr)
}
