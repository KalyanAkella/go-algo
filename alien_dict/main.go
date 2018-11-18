package main

import "fmt"

// SRC: https://practice.geeksforgeeks.org/problems/alien-dictionary/1

type Graph struct {
	nodes map[string]*Node
}

func (g *Graph) AddEdge(src, dst string) {
	srcNode, dstNode := g.nodes[src], g.nodes[dst]
	if srcNode == nil || dstNode == nil {
		panic("Src or Dst node doesn't exist")
	}
	srcNode.neighbours = append(srcNode.neighbours, dstNode)
	dstNode.inDegree++
}

func (g *Graph) topologicalSort() []*Node {
	var order []*Node
	node := g.getNodeWithInDegree(0)

	for node != nil {
		order = append(order, node)
		node.inDegree--

		for _, neighbour := range node.neighbours {
			neighbour.inDegree--
		}

		node = g.getNodeWithInDegree(0)
	}
	return order
}

func (g *Graph) getNodeWithInDegree(degree int) *Node {
	for _, node := range g.nodes {
		if node.inDegree == degree {
			return node
		}
	}
	return nil
}

type Node struct {
	val        string
	neighbours []*Node
	inDegree   int
}

func sortedChars(dict []string, k int) string {
	g := NewGraph(k)
	buildRelations(g, dict)
	nodes := g.topologicalSort()
	return buildResponse(nodes)
}

func buildResponse(nodes []*Node) string {
	ans := ""
	for _, node := range nodes {
		ans = ans + node.val
	}
	return ans
}

func buildRelations(g *Graph, dict []string) {
	for i := 0; i < len(dict); i++ {
		for j := i + 1; j < len(dict); j++ {
			earlier, later := dict[i], dict[j]
			l := min(len(earlier), len(later))
			for k := 0; k < l; k++ {
				if earlier[k] != later[k] {
					g.AddEdge(string(earlier[k]), string(later[k]))
					break
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewGraph(numNodes int) *Graph {
	g := &Graph{nodes: make(map[string]*Node)}
	for i := 0; i < numNodes; i++ {
		char := string('a' + i)
		g.nodes[char] = &Node{val: char, neighbours: make([]*Node, 0)}
	}
	return g
}

func main() {
	fmt.Println(sortedChars([]string{"baa", "abcd", "abca", "cab", "cad"}, 4))
	fmt.Println(sortedChars([]string{"caa", "aaa", "aab"}, 3))
}
