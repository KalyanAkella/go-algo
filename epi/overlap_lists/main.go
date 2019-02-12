package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

func overlappingNode(h1, h2 *Node) *Node {
	var l1, l2 int

	for t1 := h1; t1 != nil; t1 = t1.next {
		l1++
	}

	for t2 := h2; t2 != nil; t2 = t2.next {
		l2++
	}

	var small, large *Node
	var diff int
	if l1 < l2 {
		small, large = h1, h2
		diff = l2 - l1
	} else {
		small, large = h2, h1
		diff = l1 - l2
	}

	for diff > 0 {
		large = large.next
	}

	for small != large {
		small, large = small.next, large.next
	}

	return small
}

func overlappingNodeiWithLists(h1, h2 *list.List) *list.Element {
	var l1, l2 int

	for e1 := h1.Front(); e1 != nil; e1 = e1.Next() {
		l1++
	}

	for e2 := h2.Front(); e2 != nil; e2 = e2.Next() {
		l2++
	}

	var small, large *list.List
	var diff int
	if l1 > l2 {
		small, large = h2, h1
		diff = l1 - l2
	} else {
		small, large = h1, h2
		diff = l2 - l1
	}

	var e *list.Element
	for e = large.Front(); diff > 0; diff, e = diff-1, e.Next() {
	}

	for e1 := small.Front(); e1 != e; e, e1 = e.Next(), e1.Next() {
	}

	return e
}

func main() {
	a := &Node{"a", nil}
	b := &Node{"b", nil}
	c := &Node{"c", nil}
	d := &Node{"d", nil}
	e := &Node{"e", nil}
	p := &Node{"p", nil}
	q := &Node{"q", nil}
	r := &Node{"r", nil}

	a.next = b
	b.next = c
	c.next = d
	d.next = e

	p.next = q
	q.next = r
	r.next = d

	node := overlappingNode(a, p)
	fmt.Println(node.val)
}
