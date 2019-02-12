package main

import "fmt"
import "container/list"

func addAllRottenOranges(orngs [][]int, q *list.List) {
	for i := range orngs {
		for j := range orngs[i] {
			orng := []int{i, j}
			if isRottenOrange(orngs, orng) {
				q.PushBack(orng)
			}
		}
	}
}

func isRottenOrange(orngs [][]int, orng []int) bool {
	x, y := orng[0], orng[1]
	return orngs[x][y] == 2
}

func isFreshOrange(orngs [][]int, orng []int) bool {
	x, y := orng[0], orng[1]
	return orngs[x][y] == 1
}

func addSentinel(q *list.List) {
	q.PushBack([]int{-1, -1})
}

func isSentinel(orng []int) bool {
	return orng[0] == -1 && orng[1] == -1
}

func isValid(orngs [][]int, pos []int) bool {
	x, y := pos[0], pos[1]
	return x >= 0 && x < len(orngs) && y >= 0 && y < len(orngs[0])
}

func getNeighbours(orngs [][]int, orng []int) [][]int {
	x, y := orng[0], orng[1]
	off := []int{-1, 0, 1}
	var result [][]int
	for i := range off {
		for j := range off {
			if i == j || off[i]*off[j] != 0 || !isValid(orngs, []int{x + off[i], y + off[j]}) {
				continue
			}
			result = append(result, []int{x + off[i], y + off[j]})
		}
	}
	return result
}

func setRotten(orngs [][]int, orng []int) {
	x, y := orng[0], orng[1]
	orngs[x][y] = 2
}

func addOrange(orng []int, q *list.List) {
	q.PushBack(orng)
}

func hasFreshOranges(orngs [][]int) bool {
	for i := range orngs {
		for j := range orngs {
			if orngs[i][j] == 1 {
				return true
			}
		}
	}
	return false
}

func minTimeToRot(orngs [][]int) int {
	cnt := 0
	q := list.New()
	addAllRottenOranges(orngs, q)
	addSentinel(q)

	for q.Len() > 0 {
		orng := q.Remove(q.Front()).([]int)
		if isSentinel(orng) {
			if q.Len() > 0 {
				cnt++
				addSentinel(q)
			}
		} else {
			for _, neighOrng := range getNeighbours(orngs, orng) {
				if isFreshOrange(orngs, neighOrng) {
					setRotten(orngs, neighOrng)
					addOrange(neighOrng, q)
				}
			}
		}
	}

	if hasFreshOranges(orngs) {
		return -1
	}
	return cnt
}

func main() {
	orngs := [][]int{
		{2, 1, 0, 2, 1},
		{1, 0, 1, 2, 1},
		{1, 0, 0, 2, 1},
	}
	fmt.Println(minTimeToRot(orngs))
}
