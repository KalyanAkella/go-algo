package main

import "fmt"

func makeChangeHelperMemo(amount int, denoms []int, index int, memo [][]int) int {
	if memo[amount][index] == 0 {
		if index >= (len(denoms) - 1) {
			memo[amount][index] = 1
		} else {
			currDenom := denoms[index]
			ans := 0
			for i := 0; i*currDenom <= amount; i++ {
				remainingAmount := amount - i*currDenom
				ans += makeChangeHelperMemo(remainingAmount, denoms, index+1, memo)
			}
			memo[amount][index] = ans
		}
	}
	return memo[amount][index]
}

func makeChangeHelper(amount int, denoms []int, index int) int {
	if index >= (len(denoms) - 1) {
		return 1 // last denom is penny - we can always give pennies for the remaining change - so 1 way
	}
	currDenom := denoms[index]
	ans := 0
	for i := 0; i*currDenom <= amount; i++ {
		remainingAmount := amount - i*currDenom
		ans += makeChangeHelper(remainingAmount, denoms, index+1)
	}
	return ans
}

func makeChange(n int) int {
	denoms := []int{25, 10, 5, 1}
	//return makeChangeHelper(n, denoms, 0)
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, len(denoms))
	}
	return makeChangeHelperMemo(n, denoms, 0, memo)
}

func main() {
	fmt.Println(makeChange(100000))
}
