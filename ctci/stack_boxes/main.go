package main

import (
	"fmt"
	"sort"
)

type Box struct {
	width, height, depth uint
}

func (this *Box) CanBeAbove(that *Box) bool {
	return this.width < that.width && this.height < that.height && this.depth < that.depth
}

type Boxes []*Box

func (bs Boxes) Len() int {
	return len(bs)
}

func (bs Boxes) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

type ByHeight struct{ Boxes }

func (s ByHeight) Less(i, j int) bool {
	return s.Boxes[i].height < s.Boxes[j].height
}

func computeHeight(boxes []*Box) uint {
	sort.Sort(ByHeight{boxes})
	maxHeight := uint(0)
	for i := range boxes {
		maxHeight = max(maxHeight, computeHeightForStack(boxes, i))
	}
	return maxHeight
}

//TODO: Memoize
func computeHeightForStack(boxes []*Box, bottomIndex int) uint {
	bottomBox := boxes[bottomIndex]
	stackHeight := uint(0)
	for boxIndex := bottomIndex + 1; boxIndex < len(boxes); boxIndex++ {
		box := boxes[boxIndex]
		if box.CanBeAbove(bottomBox) {
			height := computeHeightForStack(boxes, boxIndex)
			stackHeight = max(stackHeight, height)
		}
	}
	stackHeight += bottomBox.height
	return stackHeight
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
}
