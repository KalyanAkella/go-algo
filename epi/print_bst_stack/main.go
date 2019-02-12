package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val         int
	left, right *Node
}

func printBST(root *Node) {
	stk := list.New()

	curr := root
	for stk.Len() > 0 || curr != nil {
		if curr != nil {
			stk.PushBack(curr)
			curr = curr.left
		} else {
			top := stk.Remove(stk.Back()).(*Node)
			fmt.Println(top.val)
			curr = top.right
		}
	}
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
	printBST(root)
}
