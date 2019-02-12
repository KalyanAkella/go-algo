package main

import (
	"container/list"
	"errors"
	"fmt"
)

type IntStack struct {
	mainStk *list.List
	tempStk *list.List
}

func New() *IntStack {
	return &IntStack{mainStk: list.New(), tempStk: list.New()}
}

func (stk *IntStack) Push(num int) {
	for stk.mainStk.Back() != nil {
		top := stk.mainStk.Back().Value.(int)
		if num < top {
			stk.mainStk.PushBack(num)
			break
		} else {
			element := stk.mainStk.Remove(stk.mainStk.Back())
			stk.tempStk.PushBack(element)
		}
	}
	if stk.mainStk.Len() == 0 {
		stk.mainStk.PushBack(num)
	}
	for stk.tempStk.Len() > 0 {
		element := stk.tempStk.Remove(stk.tempStk.Back())
		stk.mainStk.PushBack(element)
	}
}

func (stk *IntStack) Pop() (int, error) {
	top := stk.mainStk.Back()
	if top == nil {
		return -1, errors.New("Empty Stack")
	} else {
		element := stk.mainStk.Remove(top)
		return element.(int), nil
	}
}

func (stk *IntStack) Peek() (int, error) {
	top := stk.mainStk.Back()
	if top == nil {
		return -1, errors.New("Empty Stack")
	} else {
		return top.Value.(int), nil
	}
}

func (stk *IntStack) IsEmpty() bool {
	return stk.mainStk.Len() == 0
}

func main() {
	stk := New()
	stk.Push(5)
	stk.Push(3)
	stk.Push(1)
	stk.Push(2)
	stk.Push(4)

	for !stk.IsEmpty() {
		fmt.Println(stk.Pop())
	}
	fmt.Println("vim-go")
}
