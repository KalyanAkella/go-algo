package main

import (
	"errors"
	"fmt"
	"math"
)

type Stack struct {
	mainArr []int
	maxArr  []int
	top     int
}

func NewStack(capacity int) *Stack {
	mainArr := make([]int, capacity)
	maxArr := make([]int, capacity)
	return &Stack{mainArr, maxArr, 0}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (stk *Stack) Push(num int) error {
	if stk.top == len(stk.mainArr) {
		return errors.New("Not enough capacity")
	}
	newMax := max(num, stk.Max())
	stk.mainArr[stk.top] = num
	stk.maxArr[stk.top] = newMax
	stk.top++
	return nil
}

func (stk *Stack) Pop() (int, error) {
	if stk.top == 0 {
		return 0, errors.New("Empty stack")
	}
	stk.top--
	return stk.mainArr[stk.top], nil
}

func (stk *Stack) Max() int {
	if stk.top == 0 {
		return math.MinInt32
	}
	return stk.maxArr[stk.top-1]
}

func main() {
	stk := NewStack(5)
	stk.Push(1)
	fmt.Println(stk.Max())
	stk.Push(2)
	fmt.Println(stk.Max())
	stk.Push(3)
	fmt.Println(stk.Max())
	stk.Push(4)
	fmt.Println(stk.Max())
	stk.Push(5)
	fmt.Println(stk.Max())
}
