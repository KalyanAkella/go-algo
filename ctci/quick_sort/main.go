package main

import "fmt"

func quickSort(arr []int, lo, hi int) {
	index := partition(arr, lo, hi)
	if lo < index-1 {
		quickSort(arr, lo, index-1)
	}
	if index < hi {
		quickSort(arr, index, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	mid := lo + (hi-lo)>>1
	pivot := arr[mid]
	for lo <= hi {
		for arr[lo] < pivot {
			lo++
		}

		for arr[hi] > pivot {
			hi--
		}

		if lo <= hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
			lo++
			hi--
		}
	}
	return lo
}

func main() {
	arr := []int{4, 1, 2, 9, 8, 3, 0, 5, 7, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
