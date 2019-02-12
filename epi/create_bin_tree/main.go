package main

import (
	"container/list"
	"fmt"
)

/*
Create a binary tree given the inorder and preorder traversals
*/

type Node struct {
	val         int
	left, right *Node
}

func find(arr []int, s, e, element int) int {
	for i := s; i <= e; i++ {
		if arr[i] == element {
			return i
		}
	}
	return -1
}

func createBinTreeHelper(pre []int, pre_s, pre_e int, in []int, in_s, in_e int) *Node {
	if in_s <= in_e && pre_s <= pre_e {
		rootIndex := find(in, in_s, in_e, pre[pre_s]) // find root's index
		leftSubTreeSize := rootIndex - in_s

		return &Node{
			val:   in[rootIndex],
			left:  createBinTreeHelper(pre, pre_s+1, pre_s+1+leftSubTreeSize, in, in_s, rootIndex-1),
			right: createBinTreeHelper(pre, pre_s+1+leftSubTreeSize, pre_e, in, rootIndex+1, in_e),
		}
	}
	return nil
}

func createBinTree(pre, in []int) *Node {
	return createBinTreeHelper(pre, 0, len(pre)-1, in, 0, len(in)-1)
}

func printTreeInOrder(node *Node) {
	stk := list.New()
	curr := node
	for stk.Len() > 0 || curr != nil {
		if curr != nil {
			stk.PushBack(curr)
			curr = curr.left
		} else {
			top := stk.Remove(stk.Back()).(*Node)
			fmt.Print(top.val, " ")
			curr = top.right
		}
	}
}

func main() {
	pre := []int{20, 15, 10, 5, 13, 25, 23, 24}
	in := []int{5, 10, 13, 15, 20, 23, 24, 25}
	root := createBinTree(pre, in)
	printTreeInOrder(root)
}
