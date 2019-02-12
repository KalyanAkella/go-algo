package main

import "fmt"

type Node struct {
	val         int
	parent      *Node
	left, right *Node
}

func getSuccessor(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.right == nil {
		q := root
		x := q.parent
		for x != nil && x.right == q {
			q = x
			q = q.parent
		}
		return q
	}

	succ := root.right
	for succ.left != nil {
		succ = succ.left
	}
	return succ
}

func createBST(parent *Node, arr []int, lo, hi int) *Node {
	if lo > hi {
		return nil
	}

	mid := lo + (hi-lo)>>1
	curr := &Node{val: arr[mid], parent: parent}
	curr.left = createBST(curr, arr, lo, mid-1)
	curr.right = createBST(curr, arr, mid+1, hi)
	return curr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	root := createBST(nil, arr, 0, len(arr)-1)
	fmt.Println(getSuccessor(getSuccessor(getSuccessor(getSuccessor(root)))))
}
