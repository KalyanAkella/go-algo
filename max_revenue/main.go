package main

import (
	"container/heap"
	"fmt"
)

type Window int

type Windows []Window

func (this Windows) Len() int {
	return len(this)
}

func (this Windows) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Windows) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this *Windows) Push(val interface{}) {
	window := val.(Window)
	*this = append(*this, window)
}

func (this *Windows) Pop() interface{} {
	old := *this
	l := len(old)
	top := old[l-1]
	*this = old[0 : l-1]
	return top
}

func maxRevenue(windows *Windows, numTickets int) int {
	heap.Init(windows)

	result := 0
	for ; numTickets > 0; numTickets-- {
		window := heap.Pop(windows).(Window)
		result += int(window)
		heap.Push(windows, window-1)
	}
	return result
}

func main() {
	windows := &Windows{5, 1, 7, 10, 11, 9}
	result := maxRevenue(windows, 4)
	fmt.Println(result)
}
