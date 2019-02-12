package main

import (
	"fmt"
	"math"
)

type Node struct {
	key         int
	left, right *Node
}

func isBST(root *Node) bool {
	return checkBST(root, math.MinInt32, math.MaxInt32)
}

func checkBST(root *Node, min, max int) bool {
	if root == nil {
		return true
	}

	return root.key >= min && root.key <= max &&
		checkBST(root.left, min, root.key) &&
		checkBST(root.right, root.key, max)
}

func createBST(arr []int, lo, hi int) *Node {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)>>1
	root := &Node{key: arr[mid]}
	root.left = createBST(arr, lo, mid-1)
	root.right = createBST(arr, mid+1, hi)
	return root
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	root := createBST(arr, 0, len(arr)-1)
	fmt.Println(isBST(root))

	root.left, root.right = root.right, root.left
	fmt.Println(isBST(root))
}
