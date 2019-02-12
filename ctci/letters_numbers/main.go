package main

import "fmt"

func longestSubArray(arr []string) []string {
	diffs := computeDiffsAtEveryIndex(arr)
	i, j := findFarthestMatchingDiff(diffs)
	return arr[(i + 1):(j + 1)]
}

func findFarthestMatchingDiff(diffs []int) (int, int) {
	lookup := make(map[int]int)
	lookup[0] = -1
	max1, max2 := 0, 0
	for i := range diffs {
		diff := diffs[i]
		if match, present := lookup[diff]; present {
			distance := i - match
			longest := max2 - max1
			if distance > longest {
				max1 = match
				max2 = i
			}
		} else {
			lookup[diff] = i
		}
	}
	return max1, max2
}

func computeDiffsAtEveryIndex(arr []string) []int {
	letters, digits := 0, 0
	delta := make([]int, len(arr))
	for i := range arr {
		c := arr[i][0]
		if c >= '0' && c <= '9' {
			digits++
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			letters++
		}
		delta[i] = letters - digits
	}
	return delta
}

func main() {
	arr := []string{"a", "1", "2", "b", "c", "d", "3", "4", "5", "6"}
	fmt.Println(longestSubArray(arr))
}
