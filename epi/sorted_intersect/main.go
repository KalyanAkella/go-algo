package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func intersectBinSearch(a, b []int) []int {
	m, n := len(a), len(b)
	var short, long []int
	if m > n {
		short = b
		long = a
	} else {
		short = a
		long = b
	}

	var result []int
	for i := 0; i < len(short); i++ {
		if i == 0 || short[i] != short[i-1] {
			num := short[i]
			lo, hi := 0, len(long)
			for lo <= hi {
				mid := lo + (hi-lo)>>1
				if long[mid] > num {
					hi = mid - 1
				} else if long[mid] < num {
					lo = mid + 1
				} else {
					result = append(result, num)
					break
				}
			}
		}
	}
	return result
}

func intersectLoop(a, b []int) []int {
	var result []int
	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] == b[j] && (i == 0 || a[i] != a[i-1]) {
			result = append(result, a[i])
			i, j = i+1, j+1
		} else if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		}
	}
	return result
}

func intersect(a, b []int) []int {
	m, n := len(a), len(b)
	if abs(m-n) > 5 {
		return intersectBinSearch(a, b)
	} else {
		return intersectLoop(a, b)
	}
}

func main() {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	b := []int{1, 1, 2, 3, 5, 8, 13}
	fmt.Println(intersectLoop(a, b))
	fmt.Println(intersectBinSearch(a, b))
}
