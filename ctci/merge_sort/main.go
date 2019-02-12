package main

import (
	"fmt"
)

func mergeSort(arr []int, temp []int, lo, hi int) {
	if lo < hi {
		mid := lo + (hi-lo)>>1
		mergeSort(arr, temp, lo, mid)
		mergeSort(arr, temp, mid+1, hi)
		merge(arr, temp, lo, mid, hi)
	}
}

func merge(arr []int, temp []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = arr[i]
	}

	p, q, i := lo, mid+1, lo
	for p <= mid && q <= hi {
		if temp[p] < temp[q] {
			arr[i] = temp[p]
			p++
		} else {
			arr[i] = temp[q]
			q++
		}
		i++
	}

	for p <= mid {
		arr[i] = temp[p]
		i++
		p++
	}
}

func merge1(arr []int, lo, mid, hi int) {
	temp := make([]int, hi-lo+1)

	p, q, i := lo, mid+1, 0
	for p <= mid && q <= hi {
		if arr[p] < arr[q] {
			temp[i] = arr[p]
			p++
		} else if arr[p] > arr[q] {
			temp[i] = arr[q]
			q++
		}
		i++
	}

	for ; p <= mid; i++ {
		temp[i] = arr[p]
		p++
	}

	for ; q <= hi; i++ {
		temp[i] = arr[q]
		q++
	}

	for k := 0; k <= (hi - lo); k++ {
		arr[k+lo] = temp[k]
	}
}

func merge2(arr []int, lo, mid, hi int) {
	temp := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		temp[i-lo] = arr[i]
	}

	p, q, i := lo-lo, mid+1-lo, 0+lo
	for p <= (mid-lo) && q <= (hi-lo) {
		if temp[p] < temp[q] {
			arr[i] = temp[p]
			p++
		} else if temp[p] > temp[q] {
			arr[i] = temp[q]
			q++
		}
		i++
	}

	for ; p <= (mid - lo); i++ {
		arr[i] = temp[p]
		p++
	}

	for ; q <= (hi - lo); i++ {
		arr[i] = temp[q]
		q++
	}
}

func main() {
	arr := []int{4, 1, 2, 9, 8, 3, 0, 5, 7, 6}
	mergeSort(arr, make([]int, len(arr)), 0, len(arr)-1)
	fmt.Println(arr)
}
