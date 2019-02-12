package main

import (
	"fmt"
	"strings"
)

func recordNumEmptyStrings(input []string) ([]string, map[string]int) {
	var nonSparseInput []string
	lookup := make(map[string]int)
	numEmpties := 0
	for i := range input {
		if input[i] == "" {
			numEmpties++
		} else {
			nonSparseInput = append(nonSparseInput, input[i])
			lookup[input[i]] = numEmpties
		}
	}
	return nonSparseInput, lookup
}

func binSearchSlow(input []string, given string, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)>>1
	comparison := strings.Compare(input[mid], given)
	if comparison > 0 {
		return binSearchSlow(input, given, lo, mid-1)
	} else if comparison < 0 {
		return binSearchSlow(input, given, mid+1, hi)
	} else {
		return mid
	}
}

func sparseSearch(input []string, given string) int {
	return binSearch(input, given, 0, len(input)-1)
}

func binSearch(input []string, given string, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)>>1
	if len(input[mid]) == 0 {
		left := mid - 1
		right := mid + 1
		for {
			if left < lo && right > hi {
				return -1
			}
			if left >= lo && len(input[left]) > 0 {
				mid = left
				break
			}
			if right <= hi && len(input[right]) > 0 {
				mid = right
				break
			}
			left--
			right++
		}
	}
	comparison := strings.Compare(input[mid], given)
	if comparison > 0 {
		return binSearch(input, given, lo, mid-1)
	} else if comparison < 0 {
		return binSearch(input, given, mid+1, hi)
	} else {
		return mid
	}
}

func sparseSearchSlow(input []string, given string) int {
	nonSparseInput, stringEmptySpacesLookup := recordNumEmptyStrings(input)
	resultIndex := binSearchSlow(nonSparseInput, given, 0, len(nonSparseInput)-1)
	if resultIndex >= 0 {
		return resultIndex + stringEmptySpacesLookup[given]
	}
	return -1
}

func main() {
	input := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	fmt.Println(sparseSearch(input, "at"))
	fmt.Println(sparseSearch(input, "ball"))
	fmt.Println(sparseSearch(input, "car"))
	fmt.Println(sparseSearch(input, "dad"))
}
