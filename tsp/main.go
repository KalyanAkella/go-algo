package main

import (
	"container/list"
	"fmt"
	"math"
)

func pathCost(cost [][]int, src, dst int) int {
	numVertices := len(cost)
	runningCost := make([]int, numVertices)
	parents := make([]int, numVertices)
	visited := make([]bool, numVertices)

	q := list.New()
	q.PushBack(src)
	visited[src] = true

	for q.Len() > 0 {
		index := q.Remove(q.Front()).(int)
		if index == dst {
			break
		} else {
			for i := 0; i < numVertices; i++ {
				if i == index || visited[i] {
					continue
				}
				parents[i] = index
				runningCost[i] = cost[index][i]
				visited[i] = true
				q.PushBack(i)
			}
		}
	}

	ans := 0
	for i := dst; i != src; { //&& i >= 0 && i <= numVertices;
		ans += runningCost[i]
		i = parents[i]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(cost [][]int) int {
	numVertices := len(cost)
	ans := math.MaxInt32
	for i := 1; i < numVertices; i++ {
		for j := i; j < numVertices; j++ {
			currCost := cost[0][i] + pathCost(cost, i, j) + cost[j][0]
			ans = min(ans, currCost)
		}
	}
	return ans
}

func main() {
	cost1 := [][]int{
		{0, 1000, 5000},
		{5000, 0, 1000},
		{1000, 5000, 0},
	}
	fmt.Println(minCost(cost1))

	cost2 := [][]int{
		{0, 111},
		{112, 0},
	}
	fmt.Println(minCost(cost2))
}
