package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

/*
Input:
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
Output: [null,null,null,null,null,null,null,5,7,5,4]
Explanation:
After making six .push operations, the stack is [5,7,5,7,4,5] from bottom to top.  Then:

pop() -> returns 5, as 5 is the most frequent.
The stack becomes [5,7,5,7,4].

pop() -> returns 7, as 5 and 7 is the most frequent, but 7 is closest to the top.
The stack becomes [5,7,5,4].

pop() -> returns 5.
The stack becomes [5,7,4].

pop() -> returns 4.
The stack becomes [5,7].
*/
type Item struct {
	val   int
	count int
	tsStk *list.List
	index int
}

type MaxHeap []*Item

func (this MaxHeap) Len() int {
	return len(this)
}

func (this MaxHeap) Less(i, j int) bool {
	firItem, secItem := this[i], this[j]
	if firItem.count != secItem.count {
		return firItem.count > secItem.count
	} else {
		firTs, secTs := firItem.tsStk.Back().Value.(int), secItem.tsStk.Back().Value.(int)
		return firTs > secTs
	}
}

func (this MaxHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *MaxHeap) Push(x interface{}) {
	item := x.(*Item)
	lenHeap := len(*this)
	item.index = lenHeap
	*this = append(*this, item)
}

func (this *MaxHeap) Pop() interface{} {
	old := *this
	lold := len(old)
	top := old[lold-1]
	top.index = -1 // for safety
	*this = old[0 : lold-1]
	return top
}

type FreqStack struct {
	itemLookup map[int]*Item
	itemHeap   MaxHeap
	timer      int
}

func Constructor() FreqStack {
	this := &FreqStack{
		itemLookup: make(map[int]*Item),
		itemHeap:   make(MaxHeap, 0),
		timer:      0,
	}
	heap.Init(&this.itemHeap)
	return *this
}

func (this *FreqStack) Push(x int) {
	this.timer++
	if item, present := this.itemLookup[x]; present {
		item.tsStk.PushBack(this.timer)
		item.count++
		heap.Fix(&this.itemHeap, item.index)
	} else {
		newItem := &Item{val: x, count: 1, tsStk: list.New()}
		newItem.tsStk.PushBack(this.timer)
		this.itemLookup[x] = newItem
		heap.Push(&this.itemHeap, newItem)
	}
}

func (this *FreqStack) Pop() int {
	item := heap.Pop(&this.itemHeap).(*Item)
	if item.count == 1 {
		delete(this.itemLookup, item.val)
	} else {
		item.count--
		item.tsStk.Remove(item.tsStk.Back())
		heap.Push(&this.itemHeap, item)
	}
	return item.val
}

func main() {
	stk := Constructor()
	obj := &stk
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
