package main

import (
	"container/list"
	"fmt"
)

func genAllSubSeqs(arr []int) *list.List {
	bits := 0
	for i := 0; i < len(arr); i++ {
		bits |= (1 << uint(i))
	}

	result := list.New()
	for bits >= 0 {
		var subSeq []int
		for i := 0; i < len(arr); i++ {
			if bits>>uint(i)&1 == 1 {
				subSeq = append(subSeq, arr[i])
			}
		}
		bits--
		result.PushBack(subSeq)
	}
	return result
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	result := genAllSubSeqs(arr)
	printList(result)
}
