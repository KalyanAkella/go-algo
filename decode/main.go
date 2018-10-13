package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
			End cases: "" -> 1 and "01" -> 0
	    Case 1: "12345" -> "1" || num_ways_recur("2345") + "12" || num_ways_recur("345")
			Case 2: "27345" -> "2" || num_ways_recur("7345") - since "27" is a non-existent mapping

			Runs in complexity O(2^n) due to recursion. Can be memoised to avoid repeated computation of subproblems
*/
func num_ways_recur(mappings map[int]string, input string, length int) int {
	if length == 0 {
		return 1
	}
	ans := 0
	start := len(input) - length
	case_1, _ := strconv.Atoi(string(input[start]))
	if _, present := mappings[case_1]; present {
		ans += num_ways_recur(mappings, input, length-1)
	}
	if length >= 2 {
		case_2, _ := strconv.Atoi(string(input[start : start+2]))
		if _, present := mappings[case_2]; present {
			ans += num_ways_recur(mappings, input, length-2)
		}
	}
	return ans
}

func num_ways(mappings map[string]int, input string) int {
	if strings.Contains(input, "0") {
		return 0
	}
	rev_mappings := make(map[int]string)
	for k, v := range mappings {
		rev_mappings[v] = k
	}
	return num_ways_recur(rev_mappings, input, len(input))
}

func num_ways_1(mappings map[string]int, input string) []string {
	var result []string
	rev_mappings := make(map[int]string)
	for k, v := range mappings {
		rev_mappings[v] = k
	}
	sub_result := ""
	for _, c := range input {
		lookup, _ := strconv.Atoi(string(c))
		if ans, present := rev_mappings[lookup]; present {
			sub_result += ans
		} else {
			return result
		}
	}
	result = append(result, sub_result)
	sub_result = ""
	var i int
	for i = 0; i < len(input)-1; i++ {
		lookup, _ := strconv.Atoi(input[i:(i + 2)])
		if ans, present := rev_mappings[lookup]; present {
			sub_result += ans
		}
	}
	lookup, _ := strconv.Atoi(input[i:len(input)])
	sub_result += rev_mappings[lookup]
	result = append(result, sub_result)
	return result
}

func main() {
	mappings := make(map[string]int)
	for i := 0; i < 26; i++ {
		mappings[string('a'+i)] = i + 1
	}
	fmt.Println(num_ways(mappings, "12"))
	fmt.Println(num_ways(mappings, "27"))
}
