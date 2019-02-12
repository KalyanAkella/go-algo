package main

import (
	"fmt"
	"math"
	"math/rand"
)

func getRandom(floor, ceiling int) int {
	return rand.Intn(ceiling-floor) + floor
}

func RandomPartition(arr []int, b, e int) {
	Partition(arr, b, e, getRandom(b, e))
}

func Partition(arr []int, p, r, xi int) {
	x := arr[xi]
	i := p - 1
	for j := p; j <= r; j++ {
		if arr[j] == x {
			continue
		}
		if arr[j] <= x {
			i := i + 1
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[xi] = arr[xi], arr[i+1]
}

func nth_order_stat(arr []int, n int) int {
	var result int
	for i := 1; i <= n; i++ {
		min_idx := 0
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[min_idx] {
				min_idx = i
			}
		}
		if i == n {
			result = arr[min_idx]
		} else {
			arr[min_idx] = math.MaxInt32
		}
	}
	return result
}

func main() {
	for i := 1; i <= 6; i++ {
		arr := []int{1, 6, 3, 9, 8, 5}
		fmt.Println(nth_order_stat(arr, i))
	}
}
