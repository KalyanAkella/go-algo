package main

import (
	"fmt"
	"math"
)

func computeMinMaxPages(books []int) (int, int) {
	min, max := math.MaxInt32, 0
	for _, book := range books {
		if book < min {
			min = book
		}
		max += book
	}
	return min, max
}

func allotMinPages(books []int, N int) int {
	min, max := computeMinMaxPages(books)

	for min < max {
		mid := min + (max-min)>>1
		if isPossible(mid, books, N) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}

func isPossible(threshold int, books []int, numStudents int) bool {
	totalPagesPerStudent := 0
	numStudentsRequired := 1
	for _, book := range books {
		if book > threshold {
			return false
		}

		if (totalPagesPerStudent + book) > threshold {
			numStudentsRequired++

			totalPagesPerStudent = book

			if numStudentsRequired > numStudents {
				return false
			}
		} else {
			totalPagesPerStudent += book
		}
	}
	return true
}

func main() {
	books := []int{12, 34, 67, 90}
	N := 2
	fmt.Println(allotMinPages(books, N))
}
