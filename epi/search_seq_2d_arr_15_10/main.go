package main

import (
	"fmt"
)

func hash(i, j, off int) int {
	return i ^ j ^ off
}

func subSeqExists(matrix [][]int, seq []int) bool {
	noMatches := make(map[int]struct{})
	for i := range matrix {
		for j := range matrix[i] {
			if subSeqHelper(matrix, seq, i, j, 0, noMatches) {
				return true
			}
		}
	}
	return false
}

func subSeqHelper(matrix [][]int, seq []int, i, j, off int, noMatches map[int]struct{}) bool {
	if off == len(seq) {
		return true
	}
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		fmt.Println("invalid index", i, j, off)
		return false
	}
	noMatchKey := hash(i, j, off)
	if _, noMatch := noMatches[noMatchKey]; noMatch {
		fmt.Println("no match", i, j, off)
		return false
	}
	if matrix[i][j] == seq[off] {
		return subSeqHelper(matrix, seq, i-1, j, off+1, noMatches) ||
			subSeqHelper(matrix, seq, i, j-1, off+1, noMatches) ||
			subSeqHelper(matrix, seq, i+1, j, off+1, noMatches) ||
			subSeqHelper(matrix, seq, i, j+1, off+1, noMatches)
	}
	noMatches[noMatchKey] = struct{}{}
	fmt.Println("proper no match", i, j, off)
	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{3, 4, 5},
		{5, 6, 7},
	}
	seq := []int{1, 3, 4, 6}
	fmt.Println(subSeqExists(matrix, seq))
}
