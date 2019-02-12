package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val         interface{}
	left, right *Node
	size        uint //TODO: Logic to maintain this field needed
}

func (node *Node) GetNthNode(n uint) *Node {
	leftSize := uint(0)
	if node.left != nil {
		leftSize = node.left.size
	}

	if n < leftSize {
		return node.left.GetNthNode(n)
	} else if n == leftSize {
		return node
	} else {
		return node.right.GetNthNode(n - (leftSize + 1))
	}
}

type BinTree struct {
	root *Node
}

func (t *BinTree) Size() uint {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

func (t *BinTree) GetRandomNode() *Node {
	if t.root == nil {
		return nil
	}
	treeSize := t.Size()
	randNum := uint(rand.Intn(int(treeSize)))
	return t.root.GetNthNode(randNum)
}

//TODO: Test
func main() {
	fmt.Println("vim-go")
}
