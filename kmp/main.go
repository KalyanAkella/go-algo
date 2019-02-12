package main

import "fmt"

func computeLps(str string) []int {
	result := make([]int, len(str))
	for i, j := 1, 0; i < len(str); i++ {
		if str[i] == str[j] {
			result[i] = j + 1
			j++
		} else {
			if j == 0 {
				result[i] = j
			} else {
				j = result[j-1]
				i--
			}
		}
	}
	return result
}

//TODO
func match(patt, text string) int {
	lps := computeLps(patt)
	for i, j := 0, 0; i < len(patt) && j < len(text); i, j = i+1, j+1 {
		if patt[i] == text[j] {
			if i == len(patt)-1 {
				return j
			}
			continue
		} else {
			if j == 0 {
			} else {
			}
		}
	}
}

func main() {
	fmt.Println(match("ABCDABD", "ABC ABCDAB ABCDABCDABDE"))
}
