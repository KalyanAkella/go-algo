package main

import "fmt"

func prod(arr []int, given int) []int {
	lookup := make(map[int]struct{}, len(arr)) // Indicates a map with empty value
	for _, n := range arr {
		if given%n == 0 {
			f := given / n
			if _, present := lookup[f]; present {
				return []int{n, f}
			} else {
				lookup[n] = struct{}{}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(prod([]int{2, 6, 1, -5, 4, 8}, -20))
	fmt.Println(prod([]int{2, 6, 1, 5, 4, 8}, 23))
}
