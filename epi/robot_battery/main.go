package main

import (
	"fmt"
	"math"
)

/*
Find the minimum battery capacity needed by a robot for completing a journey of peaks and troughs.
Let B be the number of meters robot can climb on full charge.
Let z0, z1, z2, ..., z(n-1) be the Z coordinates
So, for any i, j such that, i < j and zj - zi represents the acutal climb
For robot to never run out of charge, B >= (zj - zi)
This is the problem of finding the max. difference in an array
*/
func minimumBatteryCapacity(zs []int) int {
	minZ := math.MaxInt32
	capacity := 0
	for _, z := range zs {
		minZ = min(minZ, z)
		capacity = max(capacity, z-minZ)
	}
	return capacity
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
}
