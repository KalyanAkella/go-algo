package main

import (
	"fmt"
	"math"
)

func floor(num float64) int {
	return int(math.Floor(num))
}

func ceil(num float64) int {
	return int(math.Ceil(num))
}

func median(arr1 []int, arr2 []int) int {
	l1, l2 := len(arr1), len(arr2)
	mid1, mid2 := (arr1[floor(float64(l1)/2)]+arr1[ceil(float64(l1)/2)])/2, (arr2[floor(float64(l2)/2)]+arr2[ceil(float64(l2)/2)])/2
	return 0
}

func main() {
	arr1 := []int{1, 4, 5}
	arr2 := []int{2, 3}
	fmt.Println(median(arr1, arr2))
}
