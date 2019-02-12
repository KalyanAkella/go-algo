package main

import (
	"container/list"
	"fmt"
)

func longestValidParensLengthSlow(parens string) int {
	return longestValidParens(parens, 0, len(parens)-1)
}

func longestValidParens(parens string, begin, end int) int {
	for begin < len(parens) && parens[begin] == ')' {
		begin++
	}

	for end >= 0 && parens[end] == '(' {
		end--
	}

	ans := 0
	for i := begin; i <= end; i++ {
		for j := i + 1; j <= end; j++ {
			if isBalanced(parens, i, j) {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

func isBalanced(parens string, start, end int) bool {
	stk := list.New()
	for i := start; i <= end; i++ {
		if parens[i] == '(' {
			stk.PushBack(i)
		} else {
			if stk.Len() == 0 {
				return false
			} else {
				stk.Remove(stk.Back())
			}
		}
	}
	return stk.Len() == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printStack(stk *list.List) {
	for e := stk.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

func longestValidParensLength(parens string) int {
	stk := list.New()
	stk.PushBack(-1)
	ans := 0

	for i := range parens {
		printStack(stk)
		if parens[i] == '(' {
			stk.PushBack(i)
		} else if parens[i] == ')' {
			stk.Remove(stk.Back())
			if stk.Len() > 0 {
				top := stk.Back().Value.(int)
				ans = max(ans, i-top)
			} else {
				stk.PushBack(i)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestValidParensLength(")()())"))
	fmt.Println(longestValidParensLength("(()()"))
	fmt.Println(longestValidParensLength("("))
	fmt.Println(longestValidParensLength("(("))
	fmt.Println(longestValidParensLength("))"))
	fmt.Println(longestValidParensLength("((()"))
	fmt.Println(longestValidParensLength(")()())"))
	fmt.Println(longestValidParensLength("()(()))))"))
}
