package main

import "fmt"

/*
 regex: "ab*cd.*"
 "abcdef"
 "acdef"
 "abbbcdef"
 "acd"

matches("abcdef", "ab*cd.*g") a matches a
matches("bcdef", "b*cd.*g") b matches b*
matches("cdef", "cd.*g") c matches c
matches("def", "d.*g") d matches d
matches("ef", ".*g")

*/
func matches(regex string, input string) bool {
	if len(regex) == 0 && len(input) == 0 {
		return true
	}
	var fi, fr, sr, tr byte
	if len(regex) > 0 {
		fr = regex[0]
	}
	if len(regex) > 1 {
		sr = regex[1]
	}
	if len(regex) > 2 {
		tr = regex[2]
	}
	if len(input) > 0 {
		fi = input[0]
	}
	if sr == '*' {
		if fr == fi {
			return matches(regex, input[1:])
		} else {
			for i := 0; i < len(input); i++ {
				if input[i] == tr {
					return matches(regex[2:], input[i:])
				}
			}
			return false
		}
	} else {
		if fr == fi {
			return matches(regex[1:], input[1:])
		} else {
			return false
		}
	}
}

func main() {
	fmt.Println(matches("ab*cd.*g", "abcdefg"))
	fmt.Println(matches("ab*cd.*g", "abbbcdefg"))
	fmt.Println(matches("ab*cd.*g", "abbbcdeefg"))
	fmt.Println(matches("ab*cd.*g", "abbbcdeef"))
	fmt.Println(matches("ab*cd.*g", "acdg"))
}
