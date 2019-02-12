package main

import (
	"fmt"
	"math"
)

type Listy struct {
	arr []int
}

func NewListy(arr []int) *Listy {
	internalArr := make([]int, len(arr))
	copy(internalArr, arr)
	return &Listy{arr: internalArr}
}

func (l *Listy) elementAt(index int) int {
	if index >= len(l.arr) {
		return -1
	}
	return l.arr[index]
}

func binSearch(l *Listy, num, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)>>1
	value := l.elementAt(mid)
	if value == num {
		return mid
	}
	if value == -1 || value > num {
		return binSearch(l, num, lo, mid-1)
	}
	return binSearch(l, num, mid+1, hi)
}

func search(listy *Listy, num int) int {
	return binSearch(listy, num, 0, math.MaxInt32)
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	listy := NewListy(arr)
	for i := range arr {
		fmt.Println(search(listy, arr[i]))
	}
}
