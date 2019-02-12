/*
Program to compute nCk using the formula:
nCk = (n-1)C(k-1) + nC(k-1)
*/

package main

import (
	"fmt"
)

func choose(c, k int) int {
	tbl := make([][]int, c+1)
	for i := range tbl {
		tbl[i] = make([]int, k+1)
	}
	for i := 0; i <= c; i++ {
		tbl[i][0] = 1
	}
	for i := 1; i <= k; i++ {
		tbl[i][i] = 1
	}

	for i := 2; i <= c; i++ {
		for j := 1; j < i && j <= k; j++ {
			tbl[i][j] = tbl[i][j-1] + tbl[i-1][j-1]
		}
	}
	return tbl[c][k]
}

func main() {
	fmt.Println(choose(5, 2))
}
