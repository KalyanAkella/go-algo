package main

import (
	"fmt"
	"math"
)

/*

We define f (X, Y) as number of different corresponding bits in binary representation of X and Y. For example, f (2, 7) = 2, since binary representation of 2 and 7 are 010 and 111, respectively. The first and the third bit differ, so f (2, 7) = 2.

You are given an array of N integers, A1, A2 ,..., AN. Find sum of f(Ai, Aj) for all pairs (i, j) such that 1 ≤ i, j ≤ N. Return the answer modulo 10^9+7.
*/

func abs(num int) uint {
	if num < 0 {
		return uint(-num)
	}
	return uint(num)
}

func numBits(num uint) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func numBitsSlow(num uint) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

func numBitDiffs(num1, num2 int) int {
	num := abs(num1) ^ abs(num2)
	return numBits(num)
}

const mod = 1000000007

func totalBitDiffs(arr []int) int {
	result := 0
	for i := range arr {
		for j := range arr {
			result += numBitDiffs(arr[i], arr[j]) % mod
		}
	}
	return result
}

func main() {
	fmt.Println(totalBitDiffs([]int{1, 3, 5}))
	fmt.Println(totalBitDiffs([]int{0, math.MaxInt32}))
	fmt.Println(totalBitDiffs([]int{math.MinInt32, 0}))
	fmt.Println(totalBitDiffs([]int{math.MinInt32, 0, math.MaxInt32}))
}
