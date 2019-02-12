package main

import (
	"fmt"
	"math"
)

type Node struct {
	val         interface{}
	left, right *Node
}

func isBalanced(root *Node) bool {
	return checkHeight(root) != math.MinInt32
}

func checkHeight(node *Node) int {
	if node == nil {
		return -1 // Height of an empty tree is -1
	}

	leftHeight := checkHeight(node.left)
	if leftHeight == math.MinInt32 {
		return leftHeight
	}

	rightHeight := checkHeight(node.right)
	if rightHeight == math.MinInt32 {
		return rightHeight
	}

	heightDiff := int(math.Abs(float64(leftHeight - rightHeight)))
	if heightDiff > 1 {
		return math.MinInt32
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func main() {
	root := &Node{val: 12,
		left: &Node{val: 34,
			left:  &Node{val: 42},
			right: &Node{val: 16}},
		right: &Node{val: 9,
			left: &Node{val: 37}}}
	fmt.Println(isBalanced(root))
	root.left.left.left = &Node{val: 21, left: &Node{val: 22}}
	fmt.Println(isBalanced(root))
}
