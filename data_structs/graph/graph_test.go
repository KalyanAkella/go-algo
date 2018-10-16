package graph

import (
	"os"
	"reflect"
	"testing"
)

var graph *Graph
var root *Vertex
var vertices []*Vertex

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
	a := &Vertex{Label: "a"}
	b := &Vertex{Label: "b"}
	c := &Vertex{Label: "c"}
	d := &Vertex{Label: "d"}
	e := &Vertex{Label: "e"}
	f := &Vertex{Label: "f"}
	g := &Vertex{Label: "g"}
	root = a
	graph = New(root, false)
	graph.AddEdge(a, b)
	graph.AddEdge(a, c)
	graph.AddEdge(a, e)
	graph.AddEdge(b, d)
	graph.AddEdge(b, f)
	graph.AddEdge(c, g)
	graph.AddEdge(e, f)
	vertices = []*Vertex{a, b, c, d, e, f, g}
	os.Exit(m.Run())
}

func flushMeta() {
	for _, vertex := range vertices {
		vertex.meta = nil
	}
}

func TestDepthFirstTraversal(t *testing.T) {
	defer flushMeta()
	assert(t, "Incorrect DFS Traversal", []string{"a", "b", "d", "f", "e", "c", "g"}, graph.DepthFirstTraversal_recur(root))
}

func TestBreadthFirstTraversal(t *testing.T) {
	defer flushMeta()
	assert(t, "Incorrect BFS Traversal", []string{"a", "b", "c", "e", "d", "f", "g"}, graph.BreadthFirstTraversal(root))
}
