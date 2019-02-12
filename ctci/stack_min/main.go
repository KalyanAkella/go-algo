package main

import (
	"container/list"
	"errors"
	"fmt"
)

type MinStack struct {
	mainStack *list.List
	minStack  *list.List
}

type Ordered interface {
	Compare(other interface{}) int
}

type Int int

func (num Int) Compare(other interface{}) int {
	a := int(num)
	b := int(other.(Int))
	return a - b
}

func New() *MinStack {
	return &MinStack{mainStack: list.New(), minStack: list.New()}
}

func (stk *MinStack) Push(value Ordered) {
	minValue := stk.minStack.Front()
	if minValue == nil || value.Compare(minValue.Value.(Ordered)) < 0 {
		stk.minStack.PushFront(value)
	} else {
		stk.minStack.PushFront(minValue.Value)
	}
	stk.mainStack.PushFront(value)
}

func (stk *MinStack) Pop() (Ordered, error) {
	if stk.mainStack.Len() == 0 {
		return nil, errors.New("Empty stack")
	}
	stk.minStack.Remove(stk.minStack.Front())
	return stk.mainStack.Remove(stk.mainStack.Front()).(Ordered), nil
}

func (stk *MinStack) Min() (Ordered, error) {
	if stk.minStack.Len() == 0 {
		return nil, errors.New("Empty stack")
	}
	return stk.minStack.Front().Value.(Ordered), nil
}

func (stk *MinStack) Peek() (Ordered, error) {
	if stk.mainStack.Len() == 0 {
		return nil, errors.New("Empty stack")
	}
	return stk.mainStack.Front().Value.(Ordered), nil
}

func main() {
	minStack := New()
	for i := 10; i > 0; i-- {
		minStack.Push(Int(i))
		fmt.Println(minStack.Min())
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(minStack.Min())
		minStack.Pop()
	}
	fmt.Println("vim-go")
}
