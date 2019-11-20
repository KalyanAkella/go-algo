package main

import (
	"fmt"
	"strings"
)

/*
Given an input string and ordering string, need to return true if the ordering string is present in Input string.

input = "hello world!"
ordering = "hlo!"
result = FALSE (all Ls are not before all Os)

input = "hello world!"
ordering = "!od"
result = FALSE (the input has '!' coming after 'o' and after 'd', but the pattern needs it to come before 'o' and 'd')

input = "hello world!"
ordering = "he!"
result = TRUE

input = "aaaabbbcccc"
ordering = "ac"
result = TRUE

*/

func isOrdered(input, order string) bool {
	lookup := make(map[rune]int)
	for i, c := range input {
		lookup[c] = i
	}
	prev := -1
	for _, c := range order {
		if idx, present := lookup[c]; present {
			if prev < idx {
				prev = idx
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func isOrderedSlow(input, order string) bool {
	var idxs []int
	for i := range order {
		idx := strings.LastIndexByte(input, byte(order[i]))
		idxs = append(idxs, idx)
	}

	if len(idxs) < 2 {
		return true
	} else {
		for i := 0; i < len(idxs)-1; i++ {
			if idxs[i] > idxs[i+1] {
				return false
			}
		}
		return true
	}
}

func main() {
	fmt.Println(isOrdered("hello world!", "hlo!"))
	fmt.Println(isOrdered("hello world!", "hol!"))
}
