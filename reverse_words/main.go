package main

import (
	"fmt"
)

func toChars(input string) []byte {
	var result []byte
	for i := range input {
		result = append(result, input[i])
	}
	return result
}

func toStr(input []byte) string {
	var result string
	for _, b := range input {
		result = result + string(b)
	}
	return result
}

func reverse(chars []byte) {
	for i, j := 0, len(chars)-1; i <= j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	j := 0
	for i := range chars {
		if chars[i] == ' ' {
			for k := i - 1; j <= k; j, k = j+1, k-1 {
				chars[j], chars[k] = chars[k], chars[j]
			}
			j = i + 1
		}
	}

	for k := len(chars) - 1; j <= k; j, k = j+1, k-1 {
		chars[j], chars[k] = chars[k], chars[j]
	}
}

func main() {
	fmt.Println(toStr(toChars("hello world")))
	input := "cake pound steal"
	inputChars := toChars(input)
	reverse(inputChars)
	fmt.Println(toStr(inputChars))
}
