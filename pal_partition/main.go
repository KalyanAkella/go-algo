package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numPalPartitions(input string, i, j int) int {
	if i == j {
		return 0
	}

	if isPalindrome(input[i:j]) {
		return 0
	}

	minCuts := math.MaxInt32
	for k := i + 1; k < j; k++ {
		leftPart := numPalPartitions(input, i, k)
		rightPart := numPalPartitions(input, k+1, j)
		minParts := min(leftPart, rightPart)
		minCuts = min(minCuts, 1+minParts)
	}
	return minCuts
}

func isPalindrome(input string) bool {
	for i, j := 0, len(input)-1; i <= j; i, j = i+1, j-1 {
		if input[i] != input[j] {
			return false
		}
	}
	return true
}

func main() {
	input := "ababbbabbababa"
	fmt.Println(numPalPartitions(input, 0, len(input)-1))
}
