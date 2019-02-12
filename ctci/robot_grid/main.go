package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	r, c int
}

func (p *Point) IsGreater(other *Point) bool {
	return p.r > other.r || p.c > other.c
}

func (p *Point) MoveDown() *Point {
	return &Point{p.r + 1, p.c}
}

func (p *Point) MoveRight() *Point {
	return &Point{p.r, p.c + 1}
}

func createOfflimits(offlimitsR, offlimitsC []int) map[Point]struct{} {
	offlimits := make(map[Point]struct{}, len(offlimitsR))
	for i := range offlimitsR {
		p := Point{offlimitsR[i], offlimitsC[i]}
		offlimits[p] = struct{}{}
	}
	return offlimits
}

func createMemo(rows, cols int) [][]*list.List {
	memo := make([][]*list.List, rows+1)
	for i := range memo {
		memo[i] = make([]*list.List, cols+1)
	}
	return memo
}

func findPath(rows, cols int, offlimitsR, offlimitsC []int) *list.List {
	offlimits := createOfflimits(offlimitsR, offlimitsC)
	memo := createMemo(rows, cols)
	return findPathBetween(&Point{0, 0}, &Point{rows - 1, cols - 1}, memo, offlimits)
}

// Simple memoization. Instead of storing matching paths, just store failed points instead - that will also speed up the algo
func findPathBetween(src, dst *Point, memo [][]*list.List, offlimits map[Point]struct{}) *list.List {
	r, c := src.r, src.c
	if memo[r][c] == nil {
		if src.IsGreater(dst) {
			memo[r][c] = list.New()
		} else {
			if _, present := offlimits[*src]; present {
				memo[r][c] = list.New()
			} else {
				ans := list.New()
				ans.PushBack(src)
				if *src == *dst {
					memo[r][c] = ans
				} else {
					moveDownPath := findPathBetween(src.MoveDown(), dst, memo, offlimits)
					if moveDownPath.Len() > 0 {
						ans.PushBackList(moveDownPath)
						memo[r][c] = ans
					} else {
						moveRightPath := findPathBetween(src.MoveRight(), dst, memo, offlimits)
						if moveRightPath.Len() > 0 {
							ans.PushBackList(moveRightPath)
							memo[r][c] = ans
						} else {
							memo[r][c] = list.New()
						}
					}
				}
			}
		}
	}
	return memo[r][c]
}

func findPathBetweenOld(src, dst *Point, offlimits map[Point]struct{}) *list.List {
	if src.IsGreater(dst) {
		return nil
	}

	if _, present := offlimits[*src]; present {
		return nil
	}

	ans := list.New()
	ans.PushBack(src)
	if *src == *dst {
		return ans
	}

	moveDownPath := findPathBetweenOld(src.MoveDown(), dst, offlimits)
	if moveDownPath != nil {
		ans.PushBackList(moveDownPath)
		return ans
	}

	moveRightPath := findPathBetweenOld(src.MoveRight(), dst, offlimits)
	if moveRightPath != nil {
		ans.PushBackList(moveRightPath)
		return ans
	}

	return nil
}

func main() {
	rows := 10
	cols := 20

	offlimitsR := []int{2, 4, 6, 8, 9}
	offlimitsC := []int{1, 2, 4, 6, 8}

	path := findPath(rows, cols, offlimitsR, offlimitsC)
	for e := path.Front(); e != nil; e = e.Next() {
		point := e.Value.(*Point)
		fmt.Printf("(%d, %d) ", point.r, point.c)
	}
	fmt.Println()
}
