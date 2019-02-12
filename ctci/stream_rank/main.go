package main

import "fmt"

type RankNode struct {
	val, rank   int
	left, right *RankNode
}

type IntStream struct {
	root *RankNode
}

func (this *IntStream) getRankOf(num int) int {
	rank := 0
	temp := this.root

	for temp != nil && num != temp.val {
		if num < temp.val {
			temp = temp.left
		} else {
			rank += temp.rank
			temp = temp.right
		}
	}
	if temp != nil {
		rank += temp.rank
	}
	return rank
}

func (this *IntStream) track(num int) {
	var parent *RankNode
	for temp := this.root; temp != nil; {
		parent = temp
		if num < temp.val {
			temp.rank++
			temp = temp.left
		} else if num > temp.val {
			temp = temp.right
		} else {
			temp.rank++
			return
		}
	}
	newNode := &RankNode{val: num, rank: 1}
	if parent == nil {
		this.root = newNode
	} else if num < parent.val {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}

func main() {
	arr := []int{20, 15, 25, 10, 23, 5, 13, 24}
	is := &IntStream{}
	for _, v := range arr {
		is.track(v)
	}
	for _, v := range arr {
		fmt.Println(is.getRankOf(v))
	}
}
