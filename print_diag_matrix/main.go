package main

import "fmt"

func printDiags(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(matrix[i-j][j], " ")
		}
		fmt.Println()
	}
	for i := l - 1; i >= 1; i-- {
		for j := 1; j < l; j++ {
			fmt.Print(matrix[i-j][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	printDiags(matrix)
}
