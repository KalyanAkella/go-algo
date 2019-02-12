package main

import (
	"container/list"
	"fmt"
)

const GRID_SIZE = 8

func main() {
	results := list.New()
	eightQueen(0, make([]int, GRID_SIZE), results)
	for e := results.Front(); e != nil; e = e.Next() {
		placements := e.Value.([]int)
		fmt.Println(placements)
	}
}

func eightQueen(row int, placements []int, results *list.List) {
	if row == GRID_SIZE {
		result := make([]int, len(placements))
		copy(result, placements)
		results.PushBack(result)
	} else {
		for col := 0; col < GRID_SIZE; col++ {
			if checkValid(row, col, placements) {
				placements[row] = col
				eightQueen(row+1, placements, results)
			}
		}
	}
}

func checkValid(row1, col1 int, placements []int) bool {
	for row2 := 0; row2 < row1; row2++ {
		col2 := placements[row2]

		if col1 == col2 { // same column
			return false
		}

		if abs(col1-col2) == (row1 - row2) { // diagonal
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
