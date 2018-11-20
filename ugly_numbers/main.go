package main

import (
	"fmt"
)

func hasOnlyFactors(num int, factors ...int) bool {
	for i := 0; i < len(factors); i++ {
		factor := factors[i]
		for num > 1 && num%factor == 0 {
			num = num / factor
		}
	}
	return num == 1
}

func uglyNumber(N int) int {
	num := 1
	for i := 1; i < N; num++ {
		if hasOnlyFactors(num, 2, 3, 5) {
			fmt.Println(num)
			i++
		}
	}
	return num
}

func main() {
	fmt.Println(uglyNumber(10000))
}
