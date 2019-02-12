package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	addStk *list.List
	delStk *list.List
}

func New() *MyQueue {
	return &MyQueue{addStk: list.New(), delStk: list.New()}
}

func (q *MyQueue) Add(item interface{}) {
	q.addStk.PushBack(item)
}

func (q *MyQueue) Remove() interface{} {
	shift(q.addStk, q.delStk)
	defer shift(q.delStk, q.addStk)
	return q.delStk.Remove(q.delStk.Back())
}

func shift(src, dst *list.List) {
	for src.Len() > 0 {
		element := src.Remove(src.Back())
		dst.PushBack(element)
	}
}

func (q *MyQueue) Peek() interface{} {
	shift(q.addStk, q.delStk)
	defer shift(q.delStk, q.addStk)
	return q.delStk.Back().Value
}

func (q *MyQueue) IsEmpty() bool {
	return q.addStk.Len() == 0
}

func main() {
	fmt.Println("vim-go")
}
