package main

import (
	"fmt"
	_ "math"
)

// Returns true for odd number of bits set, else false
func computeParity(num uint64) bool {
	result := 0
	for i := uint(0); i < 64; i++ {
		if (num&(1<<i))>>i == 1 {
			result ^= 1
		}
	}
	return result == 1
}

func main() {
	input := uint64(98)
	fmt.Printf("Parity of %b: %v\n", input, computeParity(input))
	/*for i := uint64(math.MaxUint64); i >= (math.MaxUint64 >> 1); i-- {
		fmt.Printf("Parity of %b: %v\n", i, computeParity(i))
	}*/
}
