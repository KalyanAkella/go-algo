package main

import "fmt"

func countOf2s(n int) int {
	count := 0
	for ; n > 1; n-- {
		count += numberOf2s(n)
	}
	return count
}

func numberOf2s(num int) int {
	count := 0
	for ; num > 0; num = num / 10 {
		digit := num % 10
		if digit == 2 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countOf2s(25))
}
