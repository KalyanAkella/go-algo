package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	mainList *list.List
	maxList  *list.List
}

func NewQueue() *Queue {
	return &Queue{mainList: list.New(), maxList: list.New()}
}

func (q *Queue) Enqueue(val int) {
	q.mainList.PushBack(val)
	for e := q.maxList.Front(); e != nil; e = e.Next() {
		num := e.Value.(int)
		if num < val {
			e.Value = val
		} else {
			break
		}
	}
	q.maxList.PushBack(val)
}

func (q *Queue) Dequeue() (int, int) {
	if q.mainList.Len() == 0 {
		return -1, -1
	}
	return q.mainList.Remove(q.mainList.Front()).(int), q.maxList.Remove(q.maxList.Front()).(int)
}

func (q *Queue) Max() int {
	return q.maxList.Front().Value.(int)
}

func main() {
	q := NewQueue()
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(8)
	q.Enqueue(4)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
