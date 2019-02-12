package main

import "fmt"

/*
Qs:
1. Can the number be negative ?
2. What are the max values ?
3. How do we handle overflows ?
*/
func add(a, b int) int {
	if a == b {
		return a << 1
	}
	if a < b {
		return add(b, a)
	}
	for b > 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func main() {
	fmt.Println(add(5, 3))
	fmt.Println(add(5, 5))
}
