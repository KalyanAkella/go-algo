package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

func convert(arr []int, begin, end int) *Node {
	if begin > end {
		return nil
	} else if begin == end {
		return &Node{value: arr[begin]}
	} else {
		mid := (begin + end) >> 1
		left_node := convert(arr, begin, mid-1)
		right_node := convert(arr, mid+1, end)
		return &Node{value: arr[mid], left: left_node, right: right_node}
	}
}

func preorder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preorder(node.left)
		preorder(node.right)
	}
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Printf("%d ", node.value)
		inorder(node.right)
	}
}

func postorder(node *Node) {
	if node != nil {
		postorder(node.left)
		postorder(node.right)
		fmt.Printf("%d ", node.value)
	}
}

func preorder_print() {
	var arr []int
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
		preorder(convert(arr, 0, len(arr)-1))
		fmt.Println()
	}
}

func inorder_print() {
	var arr []int
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
		inorder(convert(arr, 0, len(arr)-1))
		fmt.Println()
	}
}

func postorder_print() {
	var arr []int
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
		postorder(convert(arr, 0, len(arr)-1))
		fmt.Println()
	}
}

func main() {
	fmt.Println("PreOrder:")
	preorder_print()
	fmt.Println("InOrder:")
	inorder_print()
	fmt.Println("PostOrder:")
	postorder_print()
}
