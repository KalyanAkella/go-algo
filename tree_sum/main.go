package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val      int
	children []*Node
}

func findPath(root *Node, k int) *list.List {
	result := list.New()
	if root == nil {
		return result
	} else {
		if root.val == k {
			result.PushBack([]*Node{root})
		}
		for _, child := range root.children {
			subResultWithRoot := findPath(child, k-root.val)
			for e := subResultWithRoot.Front(); e != nil; e = e.Next() {
				nodes := e.Value.([]*Node)
				result.PushBack(append([]*Node{root}, nodes...))
			}
			subResultWithoutRoot := findPath(child, k)
			result.PushBackList(subResultWithoutRoot)
		}
		return result
	}
}

func test_2() {
	nodes := make([]*Node, 11)
	nodes[0] = &Node{1, nil}
	nodes[1] = &Node{3, nil}
	nodes[2] = &Node{-1, nil}
	nodes[3] = &Node{2, nil}
	nodes[4] = &Node{1, nil}
	nodes[5] = &Node{4, nil}
	nodes[6] = &Node{5, nil}
	nodes[7] = &Node{1, nil}
	nodes[8] = &Node{1, nil}
	nodes[9] = &Node{2, nil}
	nodes[10] = &Node{6, nil}

	nodes[0].children = []*Node{nodes[1], nodes[2]}
	nodes[1].children = []*Node{nodes[3], nodes[4]}
	nodes[2].children = []*Node{nodes[5], nodes[6]}
	nodes[4].children = []*Node{nodes[7]}
	nodes[5].children = []*Node{nodes[8], nodes[9]}
	nodes[6].children = []*Node{nodes[10]}

	result := findPath(nodes[0], 5)
	printPaths(result)
}

func test_1() {
	nodes := make([]*Node, 16)
	for i := range nodes {
		nodes[i] = &Node{i + 1, nil}
	}
	nodes[0].children = []*Node{nodes[1], nodes[7]}
	nodes[1].children = []*Node{nodes[2], nodes[4]}
	nodes[2].children = []*Node{nodes[3]}
	nodes[3].children = []*Node{nodes[12], nodes[11]}
	nodes[4].children = []*Node{nodes[5], nodes[6]}
	nodes[5].children = []*Node{nodes[13]}
	nodes[7].children = []*Node{nodes[8]}
	nodes[8].children = []*Node{nodes[9], nodes[15]}
	nodes[9].children = []*Node{nodes[14]}

	result := findPath(nodes[0], 25)
	printPaths(result)
}

func printPaths(result *list.List) {
	for e := result.Front(); e != nil; e = e.Next() {
		path := e.Value.([]*Node)
		for i := range path {
			fmt.Print(path[i].val, ", ")
		}
		fmt.Println()
	}
}

func main() {
	test_1()
	fmt.Println()
	test_2()
}
