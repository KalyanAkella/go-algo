package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sqrt(x int) int {
	/*
		  sqrt(x) = y
			if square(y) = x, return y
			if square(y) < x, try between (x/2)+1 to x
			if square(y) > x, try between 1 to (x/2)-1
			return as and when the diff: ans*ans - x is increasing
	*/
	lo, hi, diff := 1, x-1, x*x
	for lo < hi {
		mid := lo + (hi-lo)>>1
		ms := mid * mid
		new_diff := ms - x
		if new_diff > diff {
			return mid
		}
		diff = new_diff
		if ms < x {
			lo = mid + 1
		} else if ms > x {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(sqrt(11))
	fmt.Println(sqrt(16))
	fmt.Println(sqrt(25))
}
