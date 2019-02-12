package main

import "fmt"

type Node struct {
	val         int
	left, right *Node
}

func (this *Node) Next(key int) *Node {
	if key < this.val {
		return this.left
	} else if key > this.val {
		return this.right
	}
	return this
}

func firstEntryLargerThan(root *Node, key int) *Node {
	temp := root
	for temp != nil && temp.val != key {
		temp = temp.Next(key)
	}
	if temp == nil || temp.right == nil {
		return nil
	}
	temp = temp.right
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

func main() {
	root := &Node{19,
		&Node{7,
			&Node{3,
				&Node{2, nil, nil},
				&Node{5, nil, nil}},
			&Node{11,
				nil,
				&Node{17,
					&Node{13, nil, nil}, nil}}},
		&Node{43, nil, nil}}
	fmt.Println(firstEntryLargerThan(root, 11))
}
