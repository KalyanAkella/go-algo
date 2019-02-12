package main

import "fmt"

type Node struct {
	val         int
	label       string
	left, right *Node
}

func search(root *Node, num int) *Node {
	var result *Node
	temp := root
	for temp != nil {
		if num < temp.val {
			temp = temp.left
		} else if num > temp.val {
			temp = temp.right
		} else {
			result = temp
			temp = temp.left
		}
	}
	return result
}

func main() {
	nodeA := &Node{108, "A", nil, nil}
	nodeB := &Node{108, "B", nil, nil}
	nodeC := &Node{-10, "C", nil, nil}
	nodeD := &Node{-14, "D", nil, nil}
	nodeE := &Node{2, "E", nil, nil}
	nodeF := &Node{108, "F", nil, nil}
	nodeG := &Node{285, "G", nil, nil}
	nodeH := &Node{243, "H", nil, nil}
	nodeI := &Node{285, "I", nil, nil}
	nodeJ := &Node{401, "J", nil, nil}

	nodeA.left, nodeA.right = nodeB, nodeG
	nodeB.left, nodeB.right = nodeC, nodeF
	nodeC.left, nodeC.right = nodeD, nodeE
	nodeG.left, nodeG.right = nodeH, nodeI
	nodeI.right = nodeJ

	fmt.Println(search(nodeA, 108))
	fmt.Println(search(nodeA, 285))
	fmt.Println(search(nodeA, 143))
}
