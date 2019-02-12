package main

import (
	"fmt"
	"math"
)

func closestPair(arr []string) (int, int) {
	lookup := make(map[string]int)
	minDistance := math.MaxInt32
	minI, minJ := -1, -1
	for i, str := range arr {
		if index, present := lookup[str]; present {
			currDistance := i - index
			if currDistance < minDistance {
				minDistance = currDistance
				minI, minJ = index, i
				lookup[str] = i
			}
		} else {
			lookup[str] = i
		}
	}
	return minI, minJ
}

func main() {
	arr := []string{"All", "work", "and", "no", "play", "makes", "for", "no", "work", "no", "fun", "and", "no", "results"}
	fmt.Println(closestPair(arr))
}
