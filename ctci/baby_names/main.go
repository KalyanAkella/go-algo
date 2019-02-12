package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	count    int
	children []*Node
	visited  bool
}

type Pair struct {
	_1, _2 string
}

type Graph struct {
	nodes map[string]*Node
}

func (this *Graph) AddEdges(pairs []*Pair) {
	for _, pair := range pairs {
		src, dst := this.nodes[pair._1], this.nodes[pair._2]
		if src == nil {
			src = &Node{count: 0, children: make([]*Node, 0)}
			this.nodes[pair._1] = src
		}
		if dst == nil {
			dst = &Node{count: 0, children: make([]*Node, 0)}
			this.nodes[pair._2] = dst
		}
		src.children = append(src.children, dst)
		dst.children = append(dst.children, src)
	}
}

func dfs(g *Graph) map[string]int {
	result := make(map[string]int)
	for key, node := range g.nodes {
		if !node.visited {
			result[key] = dfsFrom(node)
		}
	}
	return result
}

func dfsFrom(node *Node) int {
	stk := list.New()
	stk.PushBack(node)
	node.visited = true
	count := 0
	for stk.Len() > 0 {
		n := stk.Remove(stk.Back()).(*Node)
		count += n.count
		for _, c := range n.children {
			if !c.visited {
				c.visited = true
				stk.PushBack(c)
			}
		}
	}
	return count
}

func babyNames(freqs map[string]int, pairs []*Pair) map[string]int {
	g := newGraph(freqs)
	g.AddEdges(pairs)
	return dfs(g)
}

func newGraph(freqs map[string]int) *Graph {
	g := &Graph{nodes: make(map[string]*Node)}
	for babyName, freq := range freqs {
		g.nodes[babyName] = &Node{count: freq, children: make([]*Node, 0)}
	}
	return g
}

func main1() {
	freqs := map[string]int{
		"John":        15,
		"Jon":         12,
		"Chris":       13,
		"Kris":        4,
		"Christopher": 19,
	}
	pairs := []*Pair{
		&Pair{_1: "Jon", _2: "John"},
		&Pair{_1: "John", _2: "Johnny"},
		&Pair{_1: "Chris", _2: "Kris"},
		&Pair{_1: "Chris", _2: "Christopher"},
	}
	result := babyNames(freqs, pairs)
	fmt.Println(result)
}

func main() {
	freqs := map[string]int{
		"John":     10,
		"Jon":      3,
		"Davis":    2,
		"Kari":     3,
		"Johnny":   11,
		"Carlton":  8,
		"Carleton": 2,
		"Jonathan": 9,
		"Carrie":   5,
	}
	pairs := []*Pair{
		&Pair{_1: "Jonathan", _2: "John"},
		&Pair{_1: "Jon", _2: "Johnny"},
		&Pair{_1: "Johnny", _2: "John"},
		&Pair{_1: "Kari", _2: "Carrie"},
		&Pair{_1: "Carleton", _2: "Carlton"},
	}
	result := babyNames(freqs, pairs)
	fmt.Println(result)
}
