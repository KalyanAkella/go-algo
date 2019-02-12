package main

import "fmt"
import "math"

type Node struct {
	val        interface{}
	neighbours []*Node
	distances  []int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (tree *Node) Diameter() int {
	maxDistance := math.MinInt32
	for i := range tree.neighbours {
		subDistance := tree.neighbours[i].Diameter()
		maxSubDistance := max(subDistance, tree.distances[i]+subDistance)
		maxDistance = max(maxDistance, maxSubDistance)
	}
	if maxDistance == math.MinInt32 {
		return 0
	}
	return maxDistance
}

func main() {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	node4 := &Node{val: 4}
	node5 := &Node{val: 5}
	node6 := &Node{val: 6}
	node7 := &Node{val: 7}
	node8 := &Node{val: 8}
	node9 := &Node{val: 9}
	node10 := &Node{val: 10}
	node11 := &Node{val: 11}
	node12 := &Node{val: 12}
	node13 := &Node{val: 13}
	node14 := &Node{val: 14}
	node15 := &Node{val: 15}
	node16 := &Node{val: 16}

	node1.neighbours = []*Node{node2, node3, node4}
	node1.distances = []int{-7, -14, -3}

	node2.neighbours = []*Node{node5, node6}
	node2.distances = []int{4, 3}

	node4.neighbours = []*Node{node7, node8}
	node4.distances = []int{2, 1}

	node8.neighbours = []*Node{node10, node11}
	node8.distances = []int{6, 4}

	node11.neighbours = []*Node{node12, node13}
	node11.distances = []int{4, 2}

	node13.neighbours = []*Node{node14, node15, node16}
	node13.distances = []int{1, 2, 3}

	node5.neighbours = []*Node{node9}
	node5.distances = []int{-6}

	fmt.Println(node1.Diameter())
}
