package main

import (
	"fmt"
)

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func solve1(n uint) uint {
	if n == 0 || n == 1 {
		return 0
	} else if n%2 == 0 {
		return 1 + solve1(n/2)
	} else if n%3 == 0 {
		return 1 + solve1(n/3)
	} else {
		return 1 + solve1(n-1)
	}
}

func solve2(n uint) uint {
	if n == 0 || n == 1 {
		return 0
	} else if n%2 == 0 {
		return 1 + min(solve2(n/2), solve2(n-1))
	} else if n%3 == 0 {
		return 1 + min(solve2(n/3), solve2(n-1))
	} else {
		return 1 + solve2(n-1)
	}
}

var memo []uint

func solve3(n uint) uint {
	for i := uint(4); i <= n; i++ {
		if i%2 == 0 {
			memo[i] = 1 + min(memo[i/2], memo[i-1])
		} else if i%3 == 0 {
			memo[i] = 1 + min(memo[i/3], memo[i-1])
		} else {
			memo[i] = 1 + memo[i-1]
		}
	}
	return memo[n]
}

func main() {
	var n uint = 99
	memo = make([]uint, n+1)
	memo[0], memo[1] = 0, 0
	memo[2], memo[3] = 1, 1
	fmt.Println(solve1(n))
	// fmt.Println(solve2(n))
	fmt.Println(solve3(n))
}
