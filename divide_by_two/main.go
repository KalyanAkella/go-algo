package main

import "fmt"

func average(lo, hi int) int {
	return lo + (hi-lo)>>1
}

func main() {
	fmt.Println(average(0, 5), (0+5)/2)
	fmt.Println(average(0, 4), (0+4)/2)
}
