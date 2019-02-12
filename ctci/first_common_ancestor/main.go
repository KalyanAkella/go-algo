package main

import "fmt"

type Node struct {
	val         interface{}
	left, right *Node
}

func firstCommonAncestor(root, p, q *Node) *Node {
	if root == nil || root == p || root == q {
		return root
	}

	pOnLeft, pOnRight := nodeExists(root.left, p), nodeExists(root.right, p)
	qOnLeft, qOnRight := nodeExists(root.left, q), nodeExists(root.right, q)

	if pOnLeft == qOnRight || pOnRight == qOnLeft {
		return root
	} else if pOnLeft == qOnLeft {
		return firstCommonAncestor(root.left, p, q)
	} else if pOnRight == qOnRight {
		return firstCommonAncestor(root.right, p, q)
	} else {
		return nil
	}
}

func nodeExists(root, node *Node) {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}
	return nodeExists(root.left, node) || nodeExists(root.right, node)
}

func main() {
	fmt.Println("vim-go")
}
