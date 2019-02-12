package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this *IntHeap) Push(val interface{}) {
	n := val.(int)
	*this = append(*this, n)
}

func (this *IntHeap) Pop() interface{} {
	old := *this
	lold := len(old)
	top := old[lold-1]
	*this = old[0 : lold-1]
	return top
}

func fixOrder(arr []int, k int) {
	for i := 0; i < len(arr)-k; i++ {
		h := &IntHeap{}
		heap.Init(h)
		for j := i; j < (i + k); j++ {
			heap.Push(h, arr[j])
		}
		for j := i; h.Len() > 0; j++ {
			arr[j] = heap.Pop(h).(int)
		}
	}
}

func main() {
	arr := []int{1, 0, 5, 4, 3, 2, 6, 7}
	fmt.Println(arr)
	fixOrder(arr, 4)
	fmt.Println(arr)
}
