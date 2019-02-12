package main

import "fmt"

type Node struct {
	val         interface{}
	left, right *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(num int) uint {
	if num < 0 {
		return uint(-num)
	}
	return uint(num)
}

func calculateHeight(root *Node) int {
	if root == nil {
		return -1
	}
	leftHeight := calculateHeight(root.left)
	rightHeight := calculateHeight(root.right)
	return 1 + max(leftHeight, rightHeight)
}

func checkBal(root *Node, k int) *Node {
	if root == nil {
		return nil
	}

	leftBal := checkBal(root.left, k)
	if leftBal != nil {
		return leftBal
	}
	rightBal := checkBal(root.right, k)
	if rightBal != nil {
		return rightBal
	} else {
		leftHeight := calculateHeight(root.left)
		rightHeight := calculateHeight(root.right)

		if abs(leftHeight-rightHeight) >= uint(k) {
			return root
		}
	}
	return nil
}

func main() {
	nodeA := &Node{val: "A"}
	nodeB := &Node{val: "B"}
	nodeC := &Node{val: "C"}
	nodeD := &Node{val: "D"}
	nodeE := &Node{val: "E"}
	nodeF := &Node{val: "F"}
	nodeG := &Node{val: "G"}
	nodeH := &Node{val: "H"}
	nodeI := &Node{val: "I"}
	nodeJ := &Node{val: "J"}
	nodeK := &Node{val: "K"}
	nodeL := &Node{val: "L"}
	nodeM := &Node{val: "M"}
	nodeN := &Node{val: "N"}
	nodeO := &Node{val: "O"}
	nodeP := &Node{val: "P"}

	nodeA.left, nodeA.right = nodeB, nodeI
	nodeB.left, nodeB.right = nodeC, nodeF
	nodeC.left, nodeC.right = nodeD, nodeE
	nodeF.right = nodeG
	nodeG.left = nodeH

	nodeI.left, nodeI.right = nodeJ, nodeO
	nodeO.right = nodeP
	nodeJ.right = nodeK
	nodeK.left, nodeK.right = nodeL, nodeN
	nodeL.right = nodeM

	fmt.Println(calculateHeight(nodeJ))

	result := checkBal(nodeA, 3)
	fmt.Println(result)
}
