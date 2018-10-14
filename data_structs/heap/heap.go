package heap

import "fmt"

type Node interface {
	CompareWith(otherNode Node) bool
}

type Heap struct {
	nodes        []Node
	size, length int
}

func New(capacity int) *Heap {
	check(capacity > 0, "Capacity must be positive")
	nodes := make([]Node, capacity)
	return &Heap{nodes: nodes, length: capacity, size: 0}
}

func (this *Heap) Root() Node {
	return this.nodes[0]
}

func check(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func (this *Heap) heapify(index int) int {
	left_child := this.LeftChild(index)
	right_child := this.RightChild(index)
	curr_node := this.nodes[index]
	var correct_index int
	var correct_node Node
	if left_child != nil && !curr_node.CompareWith(left_child) {
		correct_index = this.leftChildIndex(index)
		correct_node = left_child
	} else {
		correct_index = index
		correct_node = curr_node
	}
	if right_child != nil && !correct_node.CompareWith(right_child) {
		correct_index = this.rightChildIndex(index)
	}
	if correct_index != index {
		this.nodes[index], this.nodes[correct_index] = this.nodes[correct_index], this.nodes[index]
		return this.heapify(correct_index)
	}
	return correct_index
}

func (this *Heap) Parent(index int) Node {
	check(index >= 0 && index < this.size, fmt.Sprintf("Index must be non-negative and less than %d", this.size))
	parent_index := index >> 1
	return this.nodes[parent_index]
}

func (this *Heap) leftChildIndex(index int) int {
	check(index >= 0 && index < this.size, fmt.Sprintf("Index must be non-negative and less than %d", this.size))
	left_child_index := index << 1
	return left_child_index
}

func (this *Heap) rightChildIndex(index int) int {
	check(index >= 0 && index < this.size, fmt.Sprintf("Index must be non-negative and less than %d", this.size))
	right_child_index := 1 + (index << 1)
	return right_child_index
}

func (this *Heap) LeftChild(index int) Node {
	left_child_index := this.leftChildIndex(index)
	if left_child_index < this.length {
		return this.nodes[left_child_index]
	} else {
		return nil
	}
}

func (this *Heap) RightChild(index int) Node {
	right_child_index := this.rightChildIndex(index)
	if right_child_index < this.length {
		return this.nodes[right_child_index]
	} else {
		return nil
	}
}

func (this *Heap) Push(node Node) int {
	check(this.size < this.length, "Not enough space to add node")
	this.nodes[0], this.nodes[this.size] = node, this.nodes[0]
	this.size++
	return this.heapify(0)
}

func (this *Heap) Pop() Node {
	check(this.size > 0, "Empty this")
	result := this.nodes[0]
	this.size--
	if this.size > 0 {
		this.nodes[0] = this.nodes[this.size]
		this.heapify(0)
	}
	return result
}
