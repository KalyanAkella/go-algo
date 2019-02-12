package main

import "fmt"

func printSpiral(arr [][]int, s, e int) {
	if s > e {
		return
	}
	if s == e {
		fmt.Print(arr[s][e], ",")
		return
	}
	for i := s; i < e; i++ {
		fmt.Print(arr[s][i], ",")
	}
	for i := s; i < e; i++ {
		fmt.Print(arr[i][e], ",")
	}
	for i := e; i > s; i-- {
		fmt.Print(arr[e][i], ",")
	}
	for i := e; i > s; i-- {
		fmt.Print(arr[i][s], ",")
	}
	printSpiral(arr, s+1, e-1)
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printSpiral(arr, 0, len(arr)-1)
	fmt.Println()

	arr = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	printSpiral(arr, 0, len(arr)-1)
}
