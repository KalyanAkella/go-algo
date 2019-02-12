package main

import (
	"container/list"
)

type Node struct {
	val                 byte
	left, right         *Node
	left_set, right_set bool
}

// (a(b(d()())(e()()))(c()(b(d()())(e()()))))
func NewTree(input string, beg int) *Node {
	stk = list.New()
	for i := range input {
		curr := input[i]
		if curr != ')' {
			if curr == '(' {
				stk.PushBack(curr)
			} else {
				stk.PushBack(&Node{val: curr})
			}
		} else {
			stk.Remove(stk.Back())
			prevNode := stk.Remove(stk.Back()).(*Node)
			//TODO
		}
	}
}

func hasDupSubArr(vals []byte) bool {
	limit := len(vals) >> 1
	for i := 2; i <= limit; i++ {
		tbl := make(map[string]struct{})
		for j := 0; j <= len(vals)-i; j++ {
			e := string(vals[j:(j + i)])
			if _, present := tbl[e]; present {
				return true
			} else {
				tbl[e] = struct{}{}
			}
		}
	}
	return false
}

func hasDupSubTree(root *Node) bool {
	if root == nil {
		return false
	}
	stk := list.New()
	stk.PushBack(root)

	var nodeVals []byte
	for stk.Len() > 0 {
		node := stk.Remove(stk.Back()).(*Node)
		nodeVals = append(nodeVals, node.val)
		if node.left != nil {
			stk.PushBack(node.left)
		}
		if node.right != nil {
			stk.PushBack(node.right)
		}
	}

	return hasDupSubArr(nodeVals)
}

func main() {
}
