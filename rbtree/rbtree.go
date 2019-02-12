package rbtree

import (
	"errors"
)

type RBTree struct {
	left   *RBTree
	right  *RBTree
	parent *RBTree

	Key   int
	Value interface{}
}

func (tree *RBTree) balance() {
	// TODO
}

func (tree *RBTree) Insert(key int, payload interface{}) {
	if key > tree.Key {
		if tree.right == nil {
			tree.right = &RBTree{Key: key, Value: payload, parent: tree}
		} else {
			tree.right.Insert(key, payload)
		}
	} else if key < tree.Key {
		if tree.left == nil {
			tree.left = &RBTree{Key: key, Value: payload, parent: tree}
		} else {
			tree.left.Insert(key, payload)
		}
	} else {
		panic(errors.New("Inserting duplicate keys is not supported"))
	}
	tree.balance()
}

/*
Returns the subtree whose keys are greater than or equal (if `inclusive` is set) to the given key
*/
func (tree *RBTree) TailTree(key int, inclusive bool) *RBTree {
	return nil
}

/*
Returns the subtree whose keys are less than or equal (if `inclusive` is set) to the given key
*/
func (tree *RBTree) HeadTree(key int, inclusive bool) *RBTree {
	return nil
}

func (tree *RBTree) GetSuccessor(key int) *RBTree {
	if key < tree.Key {
		if tree.left == nil {
			return tree
		} else {
			temp := tree.left.TreeMaximum()
			for p := temp.parent; p.right == temp && p.Key >= key; {
				temp = p
				p = temp.parent
			}
			return temp
		}
	} else if key > tree.Key {
		if tree.right == nil {
			return tree
		} else {
			temp := tree.right.TreeMinimum()
			for p := temp.parent; p.left == temp && p.Key <= key; {
				temp = p
				p = temp.parent
			}
			return temp
		}
	} else {
		return tree
	}
}

func (tree *RBTree) TreeMinimum() *RBTree {
	temp := tree
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

func (tree *RBTree) TreeMaximum() *RBTree {
	temp := tree
	for temp.right != nil {
		temp = temp.right
	}
	return temp
}
