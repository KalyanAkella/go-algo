package main

import "fmt"

/*
How big is N?
*/
func countDistinctBinStrings(n uint) uint {
	if n > 32 {
		return 0
	}
	count := uint(0)
	for num := (1 << n) - 1; num >= 0; num-- {
		if hasNoConsecutive1s(num) {
			count++
		}
	}
	return count
}

func hasNoConsecutive1s(num int) bool {
	for num > 0 {
		if num&3 == 3 {
			return false
		}
		num = num >> 1
	}
	return true
}

func main() {
	fmt.Println(countDistinctBinStrings(20))
}
