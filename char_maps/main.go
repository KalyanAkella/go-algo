package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func possibleWords(input string) *list.List {
	result := list.New()
	for i := 0; i < len(input); i++ {
		number, _ := strconv.Atoi(input[:(i + 1)])
		char := 'a' + byte(number-1)
		if char < 'a' || char > 'z' {
			continue
		}
		prefix := string(char)
		subWords := possibleWords(input[(i + 1):])
		if subWords.Len() > 0 {
			for e := subWords.Front(); e != nil; e = e.Next() {
				newWord := prefix + e.Value.(string)
				result.PushBack(newWord)
			}
		} else {
			result.PushBack(prefix)
		}
	}
	return result
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

func main() {
	printList(possibleWords("1"))
	printList(possibleWords("12"))
	printList(possibleWords("123"))
	printList(possibleWords("1234"))
	printList(possibleWords("12345"))
}
