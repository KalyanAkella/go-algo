package main

import "fmt"

type Color byte

const (
	None Color = iota
	Red
	Blue
	Green
)

type Node struct {
	val        interface{}
	color      Color
	neighbours []*Node
}

type Graph struct {
	vertices []*Node
}

func (this *Graph) ColorNodes() {
	for _, node := range this.vertices {
		takenColors := make(map[Color]struct{})
		canBeColored := true
		for _, neighbour := range node.neighbours {
			if neighbour.val == node.val {
				fmt.Println("WARNING: Loop detected. Impossible to color this node")
				canBeColored = false
				break
			}
			if neighbour.color != None {
				takenColors[neighbour.color] = struct{}{}
			}
		}

		if canBeColored {
			for i := Red; i <= Green; i++ {
				if _, taken := takenColors[i]; !taken {
					node.color = i
					break
				}
			}
		}
	}
}

func (this *Graph) PrintColors() {
	for _, node := range this.vertices {
		fmt.Printf("%v -> %v\n", node.val, node.color)
	}
	fmt.Println()
}

func main() {
	nodeA := &Node{val: "A"}
	nodeB := &Node{val: "B"}
	nodeC := &Node{val: "C"}
	nodeD := &Node{val: "D"}

	nodeA.neighbours = []*Node{nodeB, nodeC, nodeD}
	nodeB.neighbours = []*Node{nodeA, nodeD}
	nodeC.neighbours = []*Node{nodeA, nodeD}
	nodeD.neighbours = []*Node{nodeB, nodeC, nodeA}

	g := &Graph{vertices: []*Node{nodeA, nodeB, nodeC, nodeD}}
	g.PrintColors()
	g.ColorNodes()
	g.PrintColors()
}
