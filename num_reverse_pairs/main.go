package main

import "fmt"
import "container/list"

type Pair struct {
	_1, _2 int
}

func numReversePairs(arr []int) *list.List {
	b := -1
	la := len(arr)
	for i := 1; i < la; i++ {
		if arr[i] < arr[i-1] {
			b = i
			break
		}
	}
	if b == -1 {
		return nil
	}

	q := list.New()
	for i, j := 0, b; i < b && j < la; {
		if arr[i] > arr[j] {
			q.PushBack(&Pair{arr[i], arr[j]})
			j++
		} else {
			i++
			j = b
		}
	}
	return q
}

func main() {
	arr := []int{4, 6, 8, 9, 1, 5, 10, 11}
	result := numReversePairs(arr)
	for e := result.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*Pair)
		fmt.Printf("(%d, %d) ", pair._1, pair._2)
	}
	fmt.Println()
}
