package main

import "fmt"

func is_interleave(input, a, b string, i, j, k int) bool {
	if i == -1 {
		return true
	}
	result := false
	currChar := input[i]
	if j >= 0 && a[j] == currChar {
		result = is_interleave(input, a, b, i-1, j-1, k)
	}
	if !result && k >= 0 && b[k] == currChar {
		result = is_interleave(input, a, b, i-1, j, k-1)
	}
	return result
}

func isInterleave(input, a, b string) bool {
	return is_interleave(input, a, b, len(input)-1, len(a)-1, len(b)-1)
}

func main() {
	fmt.Println(isInterleave("apqbcrd", "abcd", "pqr"))
	fmt.Println(isInterleave("apqbdcr", "abcd", "pqr"))
	fmt.Println(isInterleave("XXY", "XY", "X"))
	fmt.Println(isInterleave("XXY", "YX", "X"))
	fmt.Println(isInterleave("aadbcbbcac", "aabcc", "dbbca"))
}
