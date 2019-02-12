package sorts

import "math/rand"

func insertion_sort(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)
	copy(arr, nums)
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func random(s, e int) int {
	return s + rand.Intn(e-s+1)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	in := start
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			arr[i], arr[in] = arr[in], arr[i]
			in++
		}
	}
	arr[in], arr[end] = arr[end], arr[in]
	return in
}

func quick__sort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	quick__sort(arr, start, pivot-1)
	quick__sort(arr, pivot+1, end)
}

func quick_sort(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)
	copy(arr, nums)
	quick__sort(arr, 0, l-1)
	return arr
}

func merge_sort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	pivot := l >> 1
	left := merge_sort(nums[0:pivot])
	right := merge_sort(nums[pivot:l])
	return merge(left, right)
}

func merge(left, right []int) []int {
	s := len(left) + len(right)
	arr := make([]int, s)
	i, l, r := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			arr[i] = left[l]
			l++
		} else {
			arr[i] = right[r]
			r++
		}
		i++
	}
	for l < len(left) {
		arr[i] = left[l]
		i++
		l++
	}
	for r < len(right) {
		arr[i] = right[r]
		i++
		r++
	}
	return arr
}
