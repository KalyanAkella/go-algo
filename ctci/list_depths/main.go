package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value       interface{}
	Left, Right *Node
}

func listDepths(node *Node) []*list.List {
	if node == nil {
		return nil
	}

	var result []*list.List
	current := list.New()
	current.PushBack(node)

	for current.Len() > 0 {
		result = append(result, current)
		parents := current
		current = list.New()
		for parent := parents.Front(); parent != nil; parent = parent.Next() {
			parentNode := parent.Value.(*Node)
			if parentNode.Left != nil {
				current.PushBack(parentNode.Left)
			}
			if parentNode.Right != nil {
				current.PushBack(parentNode.Right)
			}
		}
	}
	return result
}

func main() {
	root := &Node{Value: 12,
		Left: &Node{Value: 34,
			Left:  &Node{Value: 42},
			Right: &Node{Value: 16}},
		Right: &Node{Value: 9,
			Left: &Node{Value: 37}}}
	depths := listDepths(root)
	for _, depth := range depths {
		for node := depth.Front(); node != nil; node = node.Next() {
			fmt.Printf("%v ", node.Value)
		}
		fmt.Println()
	}
}
