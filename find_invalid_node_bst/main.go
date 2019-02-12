package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	id          int
	val         interface{}
	left, right *TreeNode
}

func findInvalidNode(root *TreeNode, stk *list.List) *TreeNode {
	if root == nil {
		return nil
	}

	leftInvalidNode := findInvalidNode(root.left, stk)
	if leftInvalidNode != nil {
		return leftInvalidNode
	}
	currVal := root.id
	if stk.Len() > 0 {
		topVal := stk.Back().Value.(int)
		if topVal > currVal {
			return root
		} else {
			stk.Remove(stk.Back())
		}
	}
	stk.PushBack(currVal)
	rightInvalidNode := findInvalidNode(root.right, stk)
	if rightInvalidNode != nil {
		return rightInvalidNode
	}
	return nil
}

func main() {
	_8 := &TreeNode{id: 8}
	_3 := &TreeNode{id: 3}
	_10 := &TreeNode{id: 10}
	_1 := &TreeNode{id: 1}
	_6 := &TreeNode{id: 6}
	_14 := &TreeNode{id: 14}
	_4 := &TreeNode{id: 4}
	_7 := &TreeNode{id: 7}
	_13 := &TreeNode{id: 13}

	_8.left = _3
	_8.right = _10

	_3.left = _1
	_3.right = _6

	_6.left = _4
	_6.right = _7

	_10.right = _14
	_14.left = _13

	_14.id = -14

	stk := list.New()
	result := findInvalidNode(_8, stk)
	fmt.Println(result)
}
