package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	label       interface{}
	left, right *Node
}

// FIXME: Unsuccessful attempt of iterative invert
func invert2(node *Node) *Node {
	var old_parent_node, new_parent_node *Node
	stack := list.New()
	stack.PushBack(node)
	for stack.Len() > 0 {
		aNode := stack.Remove(stack.Back()).(*Node)
		if aNode != nil {
			newNode := &Node{label: aNode.label}
			if old_parent_node != nil {
				if aNode == old_parent_node.left {
					new_parent_node.right = newNode
				} else {
					new_parent_node.left = newNode
				}
			}
			stack.PushBack(aNode.left)
			stack.PushBack(aNode.right)
			old_parent_node = aNode
			new_parent_node = newNode
		}
	}
	return new_parent_node
}

func invert(node *Node) *Node {
	if node != nil {
		left_subtree := invert(node.left)
		right_subtree := invert(node.right)
		return &Node{label: node.label, right: left_subtree, left: right_subtree}
	}
	return nil
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Printf("%v ", node.label)
		inorder(node.right)
	}
}

func convert(arr []interface{}, beg, end int) *Node {
	if beg > end {
		return nil
	} else if beg == end {
		return &Node{label: arr[beg]}
	} else {
		mid := (beg + end) >> 1
		left_subtree := convert(arr, beg, mid-1)
		right_subtree := convert(arr, mid+1, end)
		return &Node{label: arr[mid], left: left_subtree, right: right_subtree}
	}
}

func main() {
	arr := []interface{}{1, 2, 3, 4, 5, 6, 7}
	tree := convert(arr, 0, len(arr)-1)
	inorder(tree)
	fmt.Println()
	mirror := invert(tree)
	inorder(mirror)
	mirror2 := invert2(tree)
	fmt.Println()
	inorder(mirror2)
}
