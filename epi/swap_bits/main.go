package main

import "fmt"

func swapBits(x, i, j uint) uint {
	if (x>>i)&1 != (x>>j)&1 {
		x ^= (1 << i) | (1 << j)
	}
	return x
}

func main() {
	num := uint(45)
	fmt.Printf("%b\n", num)
	fmt.Printf("%b\n", swapBits(num, 2, 4))
}
