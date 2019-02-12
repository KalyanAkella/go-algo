package main

import "fmt"

// O(b)
func multiplySlow(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + multiply(a, b-1)
}

// O(logb)
func multiply(a, b int) int {
	if b == 0 {
		return 0
	}
	if b&1 == 1 { //odd
		return a + multiply(a, b-1)
	} else {
		return multiply(a<<1, b>>1)
	}
}

func main() {
	fmt.Println(multiply(5, 30000000))
}
