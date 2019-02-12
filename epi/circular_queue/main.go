package main

import "fmt"

type Queue struct {
	arr        []interface{}
	head, tail int
	Size       uint
}

func NewQueue(capacity uint) *Queue {
	arr := make([]interface{}, capacity)
	return &Queue{arr, 0, 0, 0}
}

func (this *Queue) Enqueue(val interface{}) {
	if this.isFull() {
		this.resize()
	}
	this.arr[this.tail] = val
	this.tail = (this.tail + 1) % len(this.arr)
	this.Size++
}

func (this *Queue) Dequeue() interface{} {
	if this.isEmpty() {
		panic("Nothing to dequeue")
	}
	result := this.arr[this.head]
	this.Size--
	this.head = (this.head + 1) % len(this.arr)
	return result
}

func (this *Queue) isFull() bool {
	return this.head == this.tail && this.Size > 0
}

func (this *Queue) isEmpty() bool {
	return this.head == this.tail && this.Size == 0
}

func (this *Queue) resize() {
	newArr := make([]interface{}, len(this.arr)<<1)
	oldSize := this.Size
	var i int
	for i = 0; !this.isEmpty(); i++ {
		newArr[i] = this.Dequeue()
	}
	this.head = 0
	this.tail = i
	this.Size = oldSize
	this.arr = newArr
}

func main() {
	q := NewQueue(5)
	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(q.Dequeue())
	}
}
