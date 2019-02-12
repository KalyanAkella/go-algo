package main

import "fmt"

func nextPerm(input []int) []int {
	li := len(input)
	arr := make([]int, li)
	copy(arr, input)

	// find largest k where p[k] < p[k+1]
	k := li - 2
	for ; k >= 0 && arr[k] >= arr[k+1]; k-- {
	}

	// if no such k, we are at the last permutation
	if k < 0 {
		return nil
	}

	// find largest l such that p[l] > p[k]
	l := k + 1
	for ; l < li && arr[k] < arr[l]; l++ {
	}
	l--

	// swap elements at l and k
	arr[k], arr[l] = arr[l], arr[k]

	// reverse all elements from k+1 till end
	for i, j := k+1, li-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(nextPerm([]int{1, 0, 3, 2}))
	fmt.Println(nextPerm([]int{1, 3, 0, 2}))
	fmt.Println(nextPerm([]int{3, 1, 2, 0}))
}
