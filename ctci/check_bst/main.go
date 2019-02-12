package main

import (
	"fmt"
	"math"
)

type Node struct {
	val         int
	left, right *Node
}

func checkBSTRange(root *Node, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if minVal <= root.val && root.val > maxVal {
		return checkBSTRange(root.left, minVal, root.val) && checkBSTRange(root.right, root.val, maxVal)
	}
	return false
}

func checkBST(root *Node) bool {
	return checkBSTRange(root, math.MinInt32, math.MaxInt32)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &Node{val: 42,
		left: &Node{val: 34,
			left:  &Node{val: 32},
			right: &Node{val: 39}},
		right: &Node{val: 49,
			left: &Node{val: 37}}}
	fmt.Println(checkBST(root))
}
