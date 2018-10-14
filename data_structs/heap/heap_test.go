package heap

import (
	"os"
	"reflect"
	"testing"
)

var heap *Heap
var nilNode Node

type DataNode struct {
	value int
}

/*
	This comparison function emulates a Max Heap
*/
func (node *DataNode) CompareWith(otherNode Node) bool {
	otherDataNode := otherNode.(*DataNode) // Is there a better way ?
	return node.value > otherDataNode.value
}

func assert(t *testing.T, message string, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s. Expected: %v, Actual: %v", message, expected, actual)
	}
}

func assert_panic(t *testing.T, message string, testFn func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(message)
		}
	}()
	testFn()
}

func TestMain(m *testing.M) {
	heap = New(100)
	os.Exit(m.Run())
}

func TestHeap(t *testing.T) {
	assert(t, "Root must be nil initially", nilNode, heap.Root())
	assert_panic(t, "Error expected for parent of empty heap", func() { heap.Parent(0) })
	assert_panic(t, "Error expected for left child of empty heap", func() { heap.LeftChild(0) })
	assert_panic(t, "Error expected for right child of empty heap", func() { heap.RightChild(0) })

	assert(t, "Index must be 0", 0, heap.Push(&DataNode{value: 1}))
	assert(t, "Index must be 0", 0, heap.Push(&DataNode{value: 2}))
	assert(t, "Index must be 0", 0, heap.Push(&DataNode{value: 3}))
	assert(t, "Index must be 0", 0, heap.Push(&DataNode{value: 4}))

	assert(t, "Root must be 4", &DataNode{value: 4}, heap.Root())
	assert(t, "Popped element must be 4", &DataNode{value: 4}, heap.Pop())

	assert(t, "Root must be 3", &DataNode{value: 3}, heap.Root())
	assert(t, "Popped element must be 3", &DataNode{value: 3}, heap.Pop())
}
