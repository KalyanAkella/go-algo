package main

import "fmt"
import "container/list"

type Node struct {
	val         interface{}
	left, right *Node
	next        *Node
}

type Pair struct {
	node  *Node
	level int
}

func connectSameLevel(root *Node) {
	q := list.New()
	q.PushBack(&Pair{root, 0})

	var prevLevel int
	var prev *Node
	for q.Len() > 0 {
		nodeLevel := q.Remove(q.Front()).(*Pair)
		node, level := nodeLevel.node, nodeLevel.level
		if prevLevel != level {
			prev = nil
			prevLevel = level
		}
		if node.left != nil {
			q.PushBack(&Pair{node.left, level + 1})
			if prev != nil {
				prev.next = node.left
			}
			node.left.next = node.right
		} else {
			if prev != nil {
				prev.next = node.right
			}
		}
		prev = node.right
		if node.right != nil {
			q.PushBack(&Pair{node.right, level + 1})
		}
	}
}

func printList(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v -> ", root.val)
	printList(root.next)
}

func main() {
	root := &Node{10,
		&Node{3,
			&Node{4, nil, nil, nil},
			&Node{1, nil, nil, nil}, nil},
		&Node{5,
			nil,
			&Node{2, nil, nil, nil}, nil}, nil}
	connectSameLevel(root)
	printList(root)
	fmt.Println()
	printList(root.left)
	fmt.Println()
	printList(root.left.left)
}
