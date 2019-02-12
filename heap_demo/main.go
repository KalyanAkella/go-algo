package main

import (
	"container/heap"
	"fmt"
)

type Range struct {
	min, max int
}

func NewRange(min, max int) Range {
	return Range{min, max}
}

func (this Range) Middle() int {
	return this.min + (this.max-this.min)>>1
}

func (this Range) Split(pivot int) (Range, Range) {
	first := Range{min: this.min, max: pivot - 1}
	second := Range{min: pivot + 1, max: this.max}
	return first, second
}

type RangeHeap []Range

func (this RangeHeap) Len() int {
	return len(this)
}

func (this RangeHeap) Less(i, j int) bool {
	firstLen := this[i].max - this[i].min
	secondLen := this[j].max - this[j].min
	//return firstLen < secondLen
	return firstLen > secondLen
}

func (this RangeHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *RangeHeap) Push(aRange interface{}) {
	*this = append(*this, aRange.(Range))
}

func (this *RangeHeap) Pop() interface{} {
	old := *this
	length := len(old)
	top := old[length-1]
	*this = old[0 : length-1]
	return top
}

func main() {
	numSeats := 11
	ranges := &RangeHeap{NewRange(0, numSeats-1)}
	heap.Init(ranges)

	for ; numSeats >= 0; numSeats-- {
		aRange := heap.Pop(ranges).(Range)
		selectedSeat := aRange.Middle()
		fmt.Println(selectedSeat)
		range1, range2 := aRange.Split(selectedSeat)
		heap.Push(ranges, range1)
		heap.Push(ranges, range2)
	}
}
