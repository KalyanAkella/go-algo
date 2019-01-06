package graph

import (
	"container/list"
	"reflect"
)

type Vertex struct {
	Label string
	Data  interface{}
	meta  map[string]interface{}
}

type Graph struct {
	vertices  []*Vertex
	adjacency map[string][]*Vertex
	directed  bool
}

func (v *Vertex) ensureMeta() {
	if v.meta == nil {
		v.meta = make(map[string]interface{})
	}
}

func (v *Vertex) addInDegree() {
	v.ensureMeta()
	if inDeg, hasInDeg := v.meta["in_degree"]; hasInDeg {
		v.meta["in_degree"] = inDeg.(int) + 1
	} else {
		v.meta["in_degree"] = 1
	}
}

func (v *Vertex) markVisited() {
	v.ensureMeta()
	v.meta["visited"] = struct{}{}
}

func (v *Vertex) isVisited() bool {
	v.ensureMeta()
	_, visited := v.meta["visited"]
	return visited
}

func New(vertices []*Vertex, directed bool) *Graph {
	for _, vertex := range vertices {
		vertex.ensureMeta()
		vertex.meta["in_degree"] = 0
	}
	return &Graph{vertices: vertices, adjacency: make(map[string][]*Vertex), directed: directed}
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

func (g *Graph) DepthFirstTraversal(start *Vertex) []string {
	stk := list.New()
	start.markVisited()
	stk.PushBack(start)

	var result []string
	for stk.Len() > 0 {
		v := stk.Remove(stk.Back()).(*Vertex)
		result = append(result, v.Label)
		if neighbours, present := g.adjacency[v.Label]; present {
			for _, neighbour := range neighbours {
				if !neighbour.isVisited() {
					neighbour.markVisited()
					stk.PushBack(neighbour)
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

func (g *Graph) findVertexWithMeta(metaKey string, metaVal interface{}) *Vertex {
	for _, vertex := range g.vertices {
		vertex.ensureMeta()
		if reflect.DeepEqual(vertex.meta[metaKey], metaVal) {
			return vertex
		}
	}
	return nil
}

func (g *Graph) TopologicalOrder() []string {
	for _, vertex := range g.vertices {
		vertex.meta["temp_in_deg"] = vertex.meta["in_degree"]
	}
	var result []string
	for {
		v := g.findVertexWithMeta("temp_in_deg", 0)
		if v == nil {
			break
		}
		result = append(result, v.Label)
		if neighbours, present := g.adjacency[v.Label]; present {
			for _, neighbour := range neighbours {
				neighbour.meta["temp_in_deg"] = neighbour.meta["temp_in_deg"].(int) - 1
			}
		}
		v.meta["temp_in_deg"] = -1
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
	dst.addInDegree()
	if !g.directed {
		dst_label := dst.Label
		if neighbours, present := g.adjacency[dst_label]; present {
			neighbours = append(neighbours, src)
			g.adjacency[dst_label] = neighbours
		} else {
			g.adjacency[dst_label] = []*Vertex{src}
		}
		src.addInDegree()
	}
}
