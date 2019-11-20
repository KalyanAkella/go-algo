package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	id          int
	left, right *Node
}

func mirror(root *Node) *Node {
	if root == nil {
		return nil
	}

	return &Node{
		id:    root.id,
		left:  mirror(root.right),
		right: mirror(root.left),
	}
}

func printLevels(root *Node) {
	if root == nil {
		fmt.Println("nil")
	}
	q := list.New()
	q.PushBack(root)
	q.PushBack("|")

	for q.Len() > 0 {
		switch val := q.Remove(q.Front()).(type) {
		case string:
			fmt.Println()
			if q.Len() == 0 {
				break
			} else {
				q.PushBack("|")
			}
		case *Node:
			fmt.Printf("%d ", val.id)
			if val.left != nil {
				q.PushBack(val.left)
			}
			if val.right != nil {
				q.PushBack(val.right)
			}
		}
	}
}

func main() {
	bt := &Node{
		id: 10,
		left: &Node{
			id: 9,
			left: &Node{
				id: 7,
			},
			right: &Node{
				id: 6,
			},
		},
		right: &Node{
			id: 8,
			left: &Node{
				id: 5,
			},
			right: &Node{
				id: 4,
			},
		},
	}
	printLevels(bt)
	printLevels(mirror(bt))
}
