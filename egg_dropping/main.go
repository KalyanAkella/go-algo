package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numAttemptsMemo(eggs int, floors int, memo [][]int) int {
	if memo[eggs][floors] == math.MaxInt32 {
		if eggs == 1 || floors <= 1 {
			memo[eggs][floors] = floors
		} else {
			for i := 1; i <= floors; i++ {
				broken := numAttemptsMemo(eggs-1, i-1, memo)
				notBroken := numAttemptsMemo(eggs, floors-i, memo)
				subAttempts := max(broken, notBroken)
				memo[eggs][floors] = min(memo[eggs][floors], subAttempts)
			}
			memo[eggs][floors] += 1
		}
	}
	return memo[eggs][floors]
}

func numAttempts(eggs int, floors int) int {
	if eggs == 1 || floors <= 1 {
		return floors
	}
	attempts := math.MaxInt32
	for i := 1; i <= floors; i++ {
		broken := numAttempts(eggs-1, i-1)
		notBroken := numAttempts(eggs, floors-i)
		subAttempts := max(broken, notBroken)
		attempts = min(attempts, subAttempts)
	}
	return 1 + attempts
}

func initMemo(eggs, floors int) [][]int {
	memo := make([][]int, eggs+1)
	for i := range memo {
		memo[i] = make([]int, floors+1)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt32
		}
	}
	return memo
}

func main() {
	eggs, floors := 9, 3000
	//fmt.Println(numAttempts(eggs, floors))
	memo := initMemo(eggs, floors)
	fmt.Println(numAttemptsMemo(eggs, floors, memo))
}
