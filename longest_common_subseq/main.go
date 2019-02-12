package main

import (
	"fmt"
)

func longestCommonSubSequence(x, y string) int {
	lx, ly := len(x), len(y)
	c := make([][]int, lx+1)
	for i := range c {
		c[i] = make([]int, ly+1)
	}

	for i := 0; i <= lx; i++ {
		c[i][0] = 0
	}

	for i := 0; i <= ly; i++ {
		c[0][i] = 0
	}

	for i := 1; i <= lx; i++ {
		for j := 1; j <= ly; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = 1 + c[i-1][j-1]
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
		}
	}

	printSeq(c, lx, ly, x, y)
	fmt.Println()
	return c[lx][ly]
}

func printSeq(c [][]int, i, j int, x, y string) {
	if i == 0 || j == 0 {
		return
	}
	if c[i][j] == c[i-1][j-1]+1 {
		printSeq(c, i-1, j-1, x, y)
		fmt.Print(string(x[i-1]))
	} else if c[i][j] == c[i-1][j] {
		printSeq(c, i-1, j, x, y)
	} else if c[i][j] == c[i][j-1] {
		printSeq(c, i, j-1, x, y)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonSubSequence("acbaed", "abcadf"))
	//fmt.Println(longestCommonSubSequence("acbaed", "abca"))
}
