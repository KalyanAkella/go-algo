package main

import "fmt"

type Node struct {
	val         interface{}
	left, right *Node
}

func checkSubTree(main, sub *Node) bool {
	if main == nil {
		return false
	}
	if sub == nil {
		return true
	}

	node := checkExists(main, sub)
	if node == nil {
		return false
	}
	return subtreesEqual(node, sub)
}

func subtreesEqual(t1, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else {
		return t1.val == t2.val && subtreesEqual(t1.left, t2.left) && subtreesEqual(t1.right, t2.right)
	}
}

func checkExists(root, givenNode *Node) *Node {
	if root == nil || givenNode == nil || root.val == givenNode.val {
		return root
	}

	nodeInLeft := checkExists(root.left, givenNode)
	if nodeInLeft != nil {
		return nodeInLeft
	}
	nodeInRight := checkExists(root.right, givenNode)
	if nodeInRight != nil {
		return nodeInRight
	}
	return nil
}

func main() {
	t1 := &Node{val: 4,
		left: &Node{val: 2,
			left:  &Node{val: 1},
			right: &Node{val: 3}},
		right: &Node{val: 5,
			right: &Node{val: 10,
				left:  &Node{val: 7},
				right: &Node{val: 12}}}}

	t2 := &Node{val: 5,
		right: &Node{val: 10,
			left:  &Node{val: 7},
			right: &Node{val: 12}}}

	fmt.Println(checkSubTree(t1, t2))

	t2.left = &Node{val: 4}
	fmt.Println(checkSubTree(t1, t2))
}
