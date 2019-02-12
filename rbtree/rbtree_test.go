package rbtree

import (
	"fmt"
	"reflect"
	"testing"
)

var tree *RBTree

func assert(t *testing.T, message string, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s. Expected: %v, Actual: %v", message, expected, actual)
	}
}

func TestMain(m *testing.M) {
	tree = &RBTree{Key: 15, Value: 15}
	tree.Insert(6, 6)
	tree.Insert(18, 18)
	tree.Insert(3, 3)
	tree.Insert(7, 7)
	tree.Insert(17, 17)
	tree.Insert(20, 20)
	tree.Insert(2, 2)
	tree.Insert(4, 4)
	tree.Insert(13, 13)
	tree.Insert(9, 9)
	m.Run()
}

func TestTreeMinimum(t *testing.T) {
	assert(t, "Incorrect TreeMinimum", 2, tree.TreeMinimum().Value)
	assert(t, "Incorrect TreeMinimum", 17, tree.right.TreeMinimum().Value)
	assert(t, "Incorrect TreeMinimum", 7, tree.left.right.TreeMinimum().Value)
}

func TestTreeMaximum(t *testing.T) {
	assert(t, "Incorrect TreeMaximum", 20, tree.TreeMaximum().Value)
	assert(t, "Incorrect TreeMaximum", 13, tree.left.TreeMaximum().Value)
	assert(t, "Incorrect TreeMaximum", 4, tree.left.left.TreeMaximum().Value)
}

func TestGetSuccessor(t *testing.T) {
	assert(t, fmt.Sprintf("Incorrect GetSuccessor. Input: %v", 13), 13, tree.GetSuccessor(13).Value)
	assert(t, fmt.Sprintf("Incorrect GetSuccessor. Input: %v", 10), 13, tree.GetSuccessor(10).Value)
	assert(t, fmt.Sprintf("Incorrect GetSuccessor. Input: %v", 5), 6, tree.GetSuccessor(5).Value)
	//assert(t, fmt.Sprintf("Incorrect GetSuccessor. Input: %v", 19), 20, tree.GetSuccessor(19).Value)
	assert(t, fmt.Sprintf("Incorrect GetSuccessor. Input: %v", 16), 17, tree.GetSuccessor(16).Value)
}

func TestTailTree(t *testing.T) {
	assert(t, fmt.Sprintf("Incorrect TailTree. Input: %v", 5), 6, tree.TailTree(5, false).Value)
	assert(t, fmt.Sprintf("Incorrect TailTree. Input: %v", 10), 13, tree.TailTree(10, false).Value)
	assert(t, fmt.Sprintf("Incorrect TailTree. Input: %v", 14), 15, tree.TailTree(14, false).Value)
	assert(t, fmt.Sprintf("Incorrect TailTree. Input: %v", 16), 17, tree.TailTree(16, false).Value)
	assert(t, fmt.Sprintf("Incorrect TailTree. Input: %v", 19), 18, tree.TailTree(19, false).Value)
}
