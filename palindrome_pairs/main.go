package main

import "fmt"

/*
Input: ["abcd","dcba","lls","s","sssll"]
Output: [[0,1],[1,0],[3,2],[2,4]]
Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]

Input: ["bat","tab","cat"]
Output: [[0,1],[1,0]]
Explanation: The palindromes are ["battab","tabbat"]

for a given string s1:

Case1: If s1 is a blank string, then for any string that is palindrome s2, s1+s2 and s2+s1 are palindrome.

Case 2: If s2 is the reversing string of s1, then s1+s2 and s2+s1 are palindrome.

Case 3: If s1[0:cut] is palindrome and there exists s2 is the reversing string of s1[cut+1:], then s2+s1 is palindrome.

Case 4: Similiar to case3. If s1[cut+1: ] is palindrome and there exists s2 is the reversing string of s1[0:cut], then s1+s2 is palindrome
*/

func palindromePairs(words []string) [][]int {
	tbl := make(map[string]int)
	for i, word := range words {
		tbl[word] = i
	}

	var result [][]int
	for i, word := range words {
		if len(word) == 0 {
			for k, v := range tbl {
				if isPalindrome(k) && i != v {
					result = append(result, []int{i, v})
					result = append(result, []int{v, i})
				}
			}
		} else {
			if v, present := tbl[reverse(word)]; present && v != i {
				result = append(result, []int{i, v})
			}
			for cut := 1; cut < len(word); cut++ {
				prefix, suffix := word[0:cut], word[cut:]
				rev_prefix, rev_suffix := reverse(prefix), reverse(suffix)
				if isPalindrome(prefix) {
					if v, present := tbl[rev_suffix]; present && v != i {
						result = append(result, []int{v, i})
					}
				}
				if isPalindrome(suffix) {
					if v, present := tbl[rev_prefix]; present && v != i {
						result = append(result, []int{i, v})
					}
				}
			}
		}
	}
	return result
}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i <= j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}
	return reverse(word[1:]) + string(word[0])
}

func main() {
	//fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
	//fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
	fmt.Println(palindromePairs([]string{"a", "abc", "aba", ""}))
	fmt.Println("[[0,3],[3,0],[2,3],[3,2]]")
}
