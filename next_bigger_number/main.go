package main

import (
	"fmt"
)

func digits(num int) []int {
	var result []int
	for num > 0 {
		digit := num % 10
		result = append([]int{digit}, result...) // Ensure digits returned in correct order - most significant to least significant
		num = num / 10
	}
	return result
}

// This will work only if the digits given in order from most significant to least significant
func toNum(digits []int) int {
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	return result
}

func nextBigNum(num int) int {
	ds := digits(num)
	l := len(ds)
	for i := l - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			lsd, msd := ds[i], ds[j]
			if msd < lsd {
				ds[i], ds[j] = ds[j], ds[i]
				return toNum(ds)
			}
		}
	}
	return num
}

func main() {
	fmt.Println(toNum([]int{1, 0, 0, 0}))
	fmt.Println(nextBigNum(123456))
	fmt.Println(nextBigNum(18765))
	fmt.Println(nextBigNum(1870))
}
