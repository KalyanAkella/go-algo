package main

import "fmt"

/*
	Assumption: only alphabets (any case)
*/
func first_recurring(str string) rune {
	var result rune
	arr := make([]bool, 52)
	for _, c := range str {
		index := int(c - 'A')
		if arr[index] {
			result = c
			break
		} else {
			arr[index] = true
		}
	}
	return result
}

func main() {
	fmt.Println(string(first_recurring("abca")))
	fmt.Println(string(first_recurring("abcA")))
	fmt.Println(string(first_recurring("")))
}
