package main

import "fmt"

func matrixSearch(matrix [][]int, num int) (int, int) {
	for row, col := 0, len(matrix[0])-1; row < len(matrix) && col >= 0; {
		if matrix[row][col] == num {
			return row, col
		} else if matrix[row][col] > num {
			col--
		} else {
			row++
		}
	}
	return -1, -1
}

func main() {
	arr := [][]int{
		{15, 20, 40, 85},
		{20, 35, 80, 95},
		{30, 55, 95, 105},
		{40, 80, 100, 120},
	}
	fmt.Println(matrixSearch(arr, 55))
	fmt.Println(matrixSearch(arr, 102))
}
