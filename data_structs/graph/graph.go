package graph

import "container/list"

type Vertex struct {
	Label string
	Data  interface{}
	meta  map[string]interface{}
}

type Graph struct {
	root      *Vertex
	adjacency map[string][]*Vertex
	directed  bool
}

func (v *Vertex) markVisited() {
	if v.meta == nil {
		v.meta = make(map[string]interface{})
	}
	v.meta["visited"] = struct{}{}
}

func (v *Vertex) isVisited() bool {
	if v.meta != nil {
		_, visited := v.meta["visited"]
		return visited
	}
	return false
}

func New(root *Vertex, directed bool) *Graph {
	return &Graph{root: root, adjacency: make(map[string][]*Vertex), directed: directed}
}

func (g *Graph) BreadthFirstTraversal(start *Vertex) []string {
	var result []string
	queue := list.New()
	start.markVisited()
	queue.PushBack(start)

	for queue.Len() > 0 {
		v := queue.Remove(queue.Front()).(*Vertex)
		result = append(result, v.Label)
		if neighbours, present := g.adjacency[v.Label]; present {
			for _, neighbour := range neighbours {
				if !neighbour.isVisited() {
					neighbour.markVisited()
					queue.PushBack(neighbour)
				}
			}
		}
	}
	return result
}

func (g *Graph) DepthFirstTraversal_recur(start *Vertex) []string {
	start.markVisited()
	result := []string{start.Label}
	if neighbours, present := g.adjacency[start.Label]; present {
		for _, neighbour := range neighbours {
			if !neighbour.isVisited() {
				result = append(result, g.DepthFirstTraversal_recur(neighbour)...)
			}
		}
	}
	return result
}

func (g *Graph) AddEdge(src, dst *Vertex) {
	src_label := src.Label
	if neighbours, present := g.adjacency[src_label]; present {
		neighbours = append(neighbours, dst)
		g.adjacency[src_label] = neighbours
	} else {
		g.adjacency[src_label] = []*Vertex{dst}
	}
	if !g.directed {
		dst_label := dst.Label
		if neighbours, present := g.adjacency[dst_label]; present {
			neighbours = append(neighbours, src)
			g.adjacency[dst_label] = neighbours
		} else {
			g.adjacency[dst_label] = []*Vertex{src}
		}
	}
}
