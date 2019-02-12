package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func pathExists(root *TreeNode, sum int) bool {
	if sum == 0 {
		return true
	}

	if root == nil {
		return false
	}

	diff := sum - root.val
	return diff == 0 || pathExists(root.left, diff) || pathExists(root.right, diff)
}

// TODO: Test
func main() {
	fmt.Println()
}
