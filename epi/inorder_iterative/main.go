package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val                 int
	parent, left, right *Node
}

func inOrder(root *Node) {
	stk := list.New()
	stk.PushBack(root)

	for stk.Len() > 0 {
		node := stk.Remove(stk.Back()).(*Node)
		if node.left != nil {
			stk.PushBack(node.left)
		} else {
			fmt.Println(node.val)
			if node.right != nil {
				stk.PushBack(node.right)
			} else {
				if node.parent != nil && node.parent.left == node && node.parent.right != nil {
					fmt.Println(node.parent.val)
					stk.PushBack(node.parent.right)
				}
			}
		}
	}
}

func main() {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	node4 := &Node{val: 4}
	node5 := &Node{val: 5}
	node6 := &Node{val: 6}

	node1.left = node2
	node1.right = node3
	node2.parent = node1
	node3.parent = node1

	node3.right = node4
	node4.parent = node3

	node4.left = node5
	node4.right = node6
	node5.parent = node4
	node6.parent = node4

	inOrder(node1)
}
