package main

import "fmt"

/*
Remove Alternate Duplicate characters from a char array you have to do it in Place.Like keeping only the odd occurences of each character.
Time complexity: O(n), Space complexity: O(1)
*/
func removeDupChars(input string) string {
	output := ""
	lookup := make([]bool, 256)
	for i := range input {
		c := input[i]
		if lookup[c] {
			lookup[c] = false
		} else {
			lookup[c] = true
			output = fmt.Sprintf("%s%s", output, string(c))
		}
	}
	return output
}

func main() {
	fmt.Println(removeDupChars("you got beautiful eyes"))
}
