package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val         interface{}
	left, right *Node
	visited     bool
}

func (this *Node) visit() {
	this.visited = true
	fmt.Println(this.val)
}

func (this *Node) isVisited() bool {
	return this.visited
}

func printLevels(root *Node) {
	if root == nil {
		return
	}

	q := list.New()
	q.PushBack(root)
	root.visit()

	for q.Len() > 0 {
		node := q.Remove(q.Front()).(*Node)

		if node.left != nil && !node.left.isVisited() {
			node.left.visit()
			q.PushBack(node.left)
		}

		if node.right != nil && !node.right.isVisited() {
			node.right.visit()
			q.PushBack(node.right)
		}
	}
}

func main() {
	node1, node2, node3, node4, node5, node6 := &Node{val: 1}, &Node{val: 2}, &Node{val: 3}, &Node{val: 4}, &Node{val: 5}, &Node{val: 6}
	node1.left = node2
	node1.right = node3

	node3.right = node4

	node4.left = node5
	node4.right = node6

	printLevels(node1)
}
