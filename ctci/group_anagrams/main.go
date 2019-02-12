package main

import (
	"fmt"
)

func sorted(str string) string {
	tbl := make([]int, 256)
	for i := range str {
		tbl[str[i]]++
	}
	ans := ""
	for i, v := range tbl {
		if v > 0 {
			ans = ans + string(i)
		}
	}
	return ans
}

func groupAnagrams(strings []string) []string {
	lookup := make(map[string][]string)
	for i := range strings {
		sortedStr := sorted(strings[i])
		if _, present := lookup[sortedStr]; present {
			lookup[sortedStr] = append(lookup[sortedStr], strings[i])
		} else {
			lookup[sortedStr] = []string{strings[i]}
		}
	}
	var ans []string
	for _, vs := range lookup {
		for _, v := range vs {
			ans = append(ans, v)
		}
	}
	return ans
}

func main() {
	strings := []string{"abc", "def", "listen", "bar", "foo", "arb", "silent"}
	fmt.Println(strings)
	ans := groupAnagrams(strings)
	fmt.Println(ans)
}
