package main

import "fmt"

func main() {
	result := longestCommonSubstring("atutorialdyn", "fordynamictutorialprogr") // tutorial
	fmt.Println(result == 8)
}

func printResult(lcs [][]int, i, j int, x, y string) {
	if i == 0 || j == 0 {
		return
	}
	if lcs[i][j] == 1+lcs[i-1][j-1] {
		printResult(lcs, i-1, j-1, x, y)
		fmt.Print(string(x[i-1]))
	}
}

func longestCommonSubstring(x, y string) int {
	lx, ly := len(x), len(y)
	lcs := make([][]int, lx+1)
	for i := range lcs {
		lcs[i] = make([]int, ly+1)
	}

	var ri, rj, result int
	for i := 1; i <= lx; i++ {
		for j := 1; j <= ly; j++ {
			if x[i-1] == y[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
				if lcs[i][j] > result {
					result = lcs[i][j]
					ri, rj = i, j
				}
			}
		}
	}
	printResult(lcs, ri, rj, x, y)
	fmt.Println()
	return result
}
