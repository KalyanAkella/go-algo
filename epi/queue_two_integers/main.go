package main

import (
	"fmt"
)

type Queue struct {
	data, size uint
}

func (this *Queue) Enqueue(num byte) {
	if num < 0 || num > 9 {
		panic("Given number must lie in [0, 9]")
	}
	this.data = this.data*10 + uint(num) // TODO: check overflow
	this.size++
}

func (this *Queue) Dequeue() byte {
	if this.size == 0 {
		panic("Nothing to dequeue")
	}
	this.size--
	factor := powOf10(this.size)
	result := this.data / factor
	this.data = this.data % factor
	return byte(result)
}

func powOf10(exp uint) uint {
	result := uint(1)
	for exp > 0 {
		result *= 10
		exp--
	}
	return result
}

func main() {
	q := &Queue{0, 0}
	for i := byte(9); i > 0; i-- {
		q.Enqueue(i)
	}
	fmt.Println(q.data)
	fmt.Println(q.Dequeue())
	fmt.Println(q.data)
}
