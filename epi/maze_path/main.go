package main

import (
	"container/list"
	"fmt"
)

type Position struct {
	row, col int
	prev     *Position
	visited  bool
}

func (this *Position) neighbours(maze [][]*Position) []*Position {
	offsets := []int{-1, 0, 1}
	var result []*Position
	for i := range offsets {
		for j := range offsets {
			if offsets[i]*offsets[j] == 0 { //diagonals not allowed
				newRow, newCol := this.row+offsets[i], this.col+offsets[j]
				if newRow >= 0 && newRow < len(maze) && newCol >= 0 && newCol < len(maze[newRow]) && maze[newRow][newCol] != nil {
					result = append(result, maze[newRow][newCol])
				}
			}
		}
	}
	return result
}

func dfsPath(maze [][]*Position, src, dst *Position) {
	if src == dst {
		return
	}
	src.visited = true
	for _, neighbour := range src.neighbours(maze) {
		if !neighbour.visited {
			neighbour.visited = true
			neighbour.prev = src
			dfsPath(maze, neighbour, dst)
		}
	}
}

func bfsPath(maze [][]*Position, src, dst *Position) {
	q := list.New()
	q.PushBack(src)
	src.visited = true

	for q.Len() > 0 {
		position := q.Remove(q.Front()).(*Position)
		for _, neighbour := range position.neighbours(maze) {
			if !neighbour.visited {
				neighbour.visited = true
				neighbour.prev = position
				q.PushBack(neighbour)
			}
		}
	}
}

func printPath(position *Position) {
	if position == nil {
		return
	}
	printPath(position.prev)
	fmt.Printf("(%d, %d) ", position.row, position.col)
}

func createMazePositions(maze [][]int) [][]*Position {
	result := make([][]*Position, len(maze))
	for i := range maze {
		result[i] = make([]*Position, len(maze[i]))
		for j := range maze[i] {
			if maze[i][j] == 1 {
				result[i][j] = &Position{row: i, col: j}
			}
		}
	}
	return result
}

func main() {
	maze := [][]int{
		{0, 1, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 1, 0, 0},
		{1, 1, 1, 0, 0, 0, 1, 1, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 1, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 1},
	}

	mazePositions := createMazePositions(maze)
	src, dst := mazePositions[9][0], mazePositions[0][9]
	//bfsPath(mazePositions, src, dst)
	dfsPath(mazePositions, src, dst)
	printPath(dst)
	fmt.Println()
}
