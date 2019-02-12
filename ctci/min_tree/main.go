package main

import "fmt"

type Node struct {
	val         int
	left, right *Node
}

func createBST(arr []int, lo, hi int) *Node {
	if lo > hi {
		return nil
	}

	mid := lo + (hi-lo)>>1
	root := &Node{val: arr[mid]}
	root.left = createBST(arr, lo, mid-1)
	root.right = createBST(arr, mid+1, hi)
	return root
}

func printInOrder(root *Node) {
	if root == nil {
		return
	}
	printInOrder(root.left)
	fmt.Println(root.val)
	printInOrder(root.right)
}

func main() {
	arr := []int{1, 2, 3, 4}
	bst := createBST(arr, 0, len(arr)-1)
	printInOrder(bst)
	fmt.Println("Root:", bst.val)
}
