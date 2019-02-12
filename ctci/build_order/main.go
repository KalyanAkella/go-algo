package main

import (
	"container/list"
	"fmt"
	"math"
)

type Node struct {
	val      interface{}
	children []*Node
	inDegree uint
	status   Status
}

type Graph struct {
	nodes []*Node
}

type Status uint

const (
	NONE      Status = 0
	PARTIAL          = 1
	COMPLETED        = 2
)

func buildOrderBasedOnDFS(g *Graph) *list.List {
	stk := list.New()
	for _, node := range g.nodes {
		if node.status == NONE {
			if !doDFS(node, stk) {
				return nil
			}
		}
	}
	return stk
}

func doDFS(node *Node, stk *list.List) bool {
	if node.status == PARTIAL {
		return false
	}

	if node.status == NONE {
		node.status = PARTIAL
		for _, child := range node.children {
			if !doDFS(child, stk) {
				return false
			}
		}
		node.status = COMPLETED
		stk.PushBack(node)
	}
	return true
}

func buildOrderBasedOnInDegree(g *Graph) *list.List {
	ans := list.New()
	numVertices := len(g.nodes)
	for i := 1; i <= numVertices; i++ {
		node := findNodeWithInDegree(0, g)
		if node == nil {
			return nil
		}
		ans.PushBack(node)
		removeGivenNode(node, g)
	}
	return ans
}

func nodeExists(givenNode *Node, g *Graph) bool {
	for _, node := range g.nodes {
		if node == givenNode {
			return true
		}
	}
	return false
}

func removeGivenNode(givenNode *Node, g *Graph) {
	if nodeExists(givenNode, g) {
		for _, child := range givenNode.children {
			child.inDegree--
		}
		givenNode.inDegree = uint(math.MaxUint32)
	}
	panic("Given node does not exist in the graph")
}

func findNodeWithInDegree(deg uint, g *Graph) *Node {
	for _, node := range g.nodes {
		if node.inDegree == deg {
			return node
		}
	}
	return nil
}

func main() {
	// TODO: Verify
	fmt.Println("vim-go")
}
