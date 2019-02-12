package main

import (
	"fmt"
	"math"
)

func min(num ...int) int {
	ans := math.MaxInt32
	for _, n := range num {
		if n < ans {
			ans = n
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the minEditDistance upto lengths i, j in a, b respectively
func minEditDistance(a, b string, i, j int) int {
	if min(i, j) == 0 {
		return max(i, j)
	}

	if a[i-1] == b[j-1] {
		return minEditDistance(a, b, i-1, j-1)
	} else {
		return 1 + min(
			minEditDistance(a, b, i-1, j-1),
			minEditDistance(a, b, i-1, j),
			minEditDistance(a, b, i, j-1),
		)
	}
}

func minEditDistanceMemo(a, b string) int {
	memo := make([][]int, len(a)+1)
	for i := range memo {
		memo[i] = make([]int, len(b)+1)
	}

	for i := 1; i <= len(a); i++ {
		memo[i][0] = i
	}

	for j := 1; j <= len(b); j++ {
		memo[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = 1 + min(memo[i-1][j-1], memo[i][j-1], memo[i-1][j])
			}
		}
	}

	return memo[len(a)][len(b)]
}

func main() {
	fmt.Println(minEditDistance("cast", "best", 4, 4))
	fmt.Println(minEditDistance("c", "b", 1, 1))
	fmt.Println(minEditDistance("", "b", 0, 1))
	fmt.Println(minEditDistance("abcdefg", "bcdefg", 7, 6))

	fmt.Println()
	fmt.Println(minEditDistanceMemo("cast", "best"))
	fmt.Println(minEditDistanceMemo("c", "b"))
	fmt.Println(minEditDistanceMemo("", "b"))
	fmt.Println(minEditDistanceMemo("abcdefg", "bcdefg"))
	fmt.Println(minEditDistanceMemo("abcdefghijklmnopqrstuvwxyz", "bcdefg123456"))
	fmt.Println(minEditDistanceMemo("carthorse", "orchestra"))
	fmt.Println(minEditDistance("abcdefghijklmnopqrstuvwxyz", "bcdefg123456", 26, 12))
}
