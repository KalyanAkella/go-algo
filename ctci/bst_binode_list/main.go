package main

import "fmt"

func main() {
	root := &BiNode{4,
		&BiNode{2,
			&BiNode{1, &BiNode{0, nil, nil}, nil},
			&BiNode{3, nil, nil}},
		&BiNode{5, nil, &BiNode{6, nil, nil}}}
	result := makeList(root)
	printList(result)
}

func printBST(root *BiNode) {
	if root == nil {
		return
	}
	printBST(root.node1)
	fmt.Printf("%d -> ", root.val)
	printBST(root.node2)
}

func printList(head *BiNode) {
	if head == nil {
		return
	}
	fmt.Printf("%d -> ", head.val)
	printList(head.node2)
}

type BiNode struct {
	val          int
	node1, node2 *BiNode
}

func makeListHelper(root *BiNode) *BiNode {
	if root == nil {
		return nil
	}

	currNode := &BiNode{val: root.val}
	if left := makeListHelper(root.node1); left != nil {
		currNode.node1 = left
		left.node2 = currNode
	}
	if right := makeListHelper(root.node2); right != nil {
		currNode.node2 = right
		right.node1 = currNode
	}
	return currNode
}

func makeList(root *BiNode) *BiNode {
	result := makeListHelper(root)
	for result.node1 != nil {
		result = result.node1
	}
	return result
}
