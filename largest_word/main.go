package main

import "fmt"

/*

{"ale", "apple", "monkey", "plea"}
"abpcplea"

*/
func numDeletions(large, small string) int {
	if len(large) < len(small) {
		return -1
	}
	freq := make([]int, 26)
	for i := range large {
		index := int(large[i] - 'a')
		freq[index]++
	}

	for i := range small {
		index := int(small[i] - 'a')
		if freq[index] > 0 {
			freq[index]--
		} else {
			freq[index]++
		}
	}

	count := 0
	for _, v := range freq {
		count += v
	}
	return count
}

func longestString(dict []string, input string) string {
	longest_str, minDeletions := dict[0], numDeletions(input, dict[0])
	for i := 1; i < len(dict); i++ {
		currDeletions := numDeletions(input, dict[i])
		if currDeletions < minDeletions {
			minDeletions = currDeletions
			longest_str = dict[i]
		}
	}
	return longest_str
}

func main() {
	dict := []string{"ale", "apple", "monkey", "plea"}
	fmt.Println(longestString(dict, "abpcplea"))

	dict = []string{"pintu", "geeksfor", "geeksgeeks", "forgeek"}
	fmt.Println(longestString(dict, "geeksforgeeks"))
}
