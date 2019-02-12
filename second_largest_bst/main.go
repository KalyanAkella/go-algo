package main

import "fmt"

type Node struct {
	val         int
	left, right *Node
}

func largestNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	result := root
	for result.right != nil {
		result = result.right
	}
	return result
}

func secondLargestNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return largestNode(root.left)
	} else {
		prev, root := root, root.right
		for root != nil {
			if root.right == nil {
				result := largestNode(root.left)
				if result != nil {
					return result
				} else {
					return prev
				}
			} else {
				prev, root = root, root.right
			}
		}
		return prev
	}
}

func main() {
	node7 := &Node{val: 7}
	node6 := &Node{val: 6}
	node15 := &Node{val: 15}
	//node16 := &Node{val: 16}
	node13 := &Node{val: 13}
	node14 := &Node{val: 14}
	node7.left, node7.right = node6, node15
	node15.left = node13
	//node15.left, node15.right = node13, node16
	node13.right = node14
	fmt.Println(secondLargestNode(node7))
}
