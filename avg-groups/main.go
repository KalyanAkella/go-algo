package main

import (
	"flag"
	"fmt"
)

var count int

func init() {
	flag.IntVar(&count, "count", count, "group count")
	flag.Parse()
}

func main() {
	count = 3
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(arr)
	i := 0
	for ; i <= (n - count); i += count {
		sum := 0
		for j := i; j < (i + count); j++ {
			sum += arr[j]
		}
		fmt.Println(sum / count)
	}

	sum := 0
	for ; i < n; i++ {
		sum += arr[i]
	}
	if sum > 0 {
		fmt.Println(sum / count)
	}
}
