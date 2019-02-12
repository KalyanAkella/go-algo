package main

import "fmt"

func closestNumWithSameWeight(num uint64) uint64 {
	for i := uint64(0); i < uint64(63); i++ {
		lsb := (num >> i) & uint64(1)
		nextLsb := (num >> (i + uint64(1))) & uint64(1)
		if lsb^nextLsb == 1 {
			return num ^ (uint64(1) << i) | (uint64(1) << (i + 1))
		}
	}
	panic("invalid input")
}

func numBits(num uint64) uint8 {
	count := uint8(0)
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func main() {
	input := uint64(15)
	output := closestNumWithSameWeight(input)
	fmt.Println(input, output, numBits(input), numBits(output))
}
