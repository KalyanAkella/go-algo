package main

import "fmt"

type Sudoku struct {
	board [][]int
}

func check(arr []int) bool {
	la := len(arr)
	if la == 9 {
		freqs := make([]int, la)
		for _, n := range arr {
			if n < 0 || n > 9 {
				return false
			}
			if n > 0 {
				freqs[n-1]++
				if freqs[n-1] > 1 {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (this *Sudoku) isValid() bool {
	for i := range this.board {
		currRow := this.board[i]
		currCol := make([]int, len(currRow))
		for j := range this.board {
			currCol[j] = this.board[j][i]
		}
		if !check(currRow) || !check(currCol) {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ro, co := i*3, j*3 //grid
			grid := make([]int, 9)
			k := 0
			for r := ro; r < ro+3; r++ {
				for c := co; c < co+3; c++ {
					grid[k] = this.board[r][c]
					k++
				}
			}
			if !check(grid) {
				return false
			}
		}
	}
	return true
}

func main() {
	sudoku := &Sudoku{board: [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}}

	fmt.Println(sudoku.isValid())
}
