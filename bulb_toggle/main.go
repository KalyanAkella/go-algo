package main

import "fmt"

func finalState(numBulbs int, lbulbs, rbulbs []int) []uint8 {
	result := make([]uint8, numBulbs)
	for i := range result {
		result[i] = uint8(1)
	}

	for i := range lbulbs {
		for j := lbulbs[i]; j <= rbulbs[i]; j++ {
			result[j-1] ^= 1
		}
	}
	return result
}

func main() {
	result := finalState(5, []int{1, 2, 3}, []int{5, 3, 5})
	fmt.Println(result)
}
