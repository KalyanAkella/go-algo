package main

import (
	"fmt"
)

type IntStack struct {
	arr          []int
	length, size int
}

func NewStack(length int) *IntStack {
	return &IntStack{arr: make([]int, length), length: length, size: 0}
}

func (stk *IntStack) Push(num int) {
	if stk.size == stk.length {
		panic("Not enough memory")
	}
	stk.arr[stk.size] = num
	stk.size++
}

func (stk *IntStack) Pop() int {
	if stk.size == 0 {
		panic("Nothing to pop")
	}
	stk.size--
	return stk.arr[stk.size]
}

func (stk *IntStack) Empty() bool {
	return stk.size == 0
}

func min_num(pattern string) int {
	stk := NewStack(9)
	result := 0
	for i := 0; i <= len(pattern); i++ {
		stk.Push(i + 1)
		if i == len(pattern) || pattern[i] == 'I' {
			for !stk.Empty() {
				result = result*10 + stk.Pop()
			}
		}
	}
	return result
}

func main() {
	fmt.Println(min_num("I"))
	fmt.Println(min_num("D"))
	fmt.Println(min_num("DD"))
	fmt.Println(min_num("II"))
	fmt.Println(min_num("DIDI"))
	fmt.Println(min_num("IIDDD"))
	fmt.Println(min_num("DDIDDIID"))
}
