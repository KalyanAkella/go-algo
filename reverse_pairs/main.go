package main

import "fmt"

func merge(arr []int, l, r int) int {
	if l > r {
		return 0
	}
	m := l + (r-l)>>1
	res := merge(arr, l, m) + merge(m+1, r)
	i, j, k, p := l, m+1, m+1, l
	for i <= m {
		for ; k <= r && arr[i] > 2*arr[k]; k++ {
		}
		res += k - (m + 1)
	}
}

func main() {
	fmt.Println(merge([]int{2, 4, 3, 5, 1}, 0, 4))
}
