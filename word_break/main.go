package main

import (
	"container/list"
	"fmt"
)

func scanSentences(input string, dict map[string]struct{}) *list.List {
	ans := list.New()
	for i := 1; i <= len(input); i++ {
		prefix := input[:i]
		if _, present := dict[prefix]; present {
			subSentences := scanSentences(input[i:], dict)
			if subSentences.Len() > 0 {
				for e := subSentences.Front(); e != nil; e = e.Next() {
					subSentence := e.Value.(string)
					newSentence := fmt.Sprintf("%s %s", prefix, subSentence)
					ans.PushBack(newSentence)
				}
			} else {
				ans.PushBack(prefix)
			}
		}
	}
	return ans
}

func printList(givenList *list.List) {
	for e := givenList.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v, ", e.Value)
	}
	fmt.Println()
}

func example1() {
	dict := map[string]struct{}{
		"snake":  struct{}{},
		"snakes": struct{}{},
		"and":    struct{}{},
		"sand":   struct{}{},
		"ladder": struct{}{},
	}
	word := "snakesandladder"
	result := scanSentences(word, dict)
	printList(result)
}

func example2() {
	dict := map[string]struct{}{
		"lr":    struct{}{},
		"m":     struct{}{},
		"lrm":   struct{}{},
		"hcdar": struct{}{},
		"wk":    struct{}{},
	}
	word := "hcdarlrm"
	result := scanSentences(word, dict)
	printList(result)
}

func main() {
	example1()
	example2()
}
