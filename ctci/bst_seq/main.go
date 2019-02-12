package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val         interface{}
	left, right *Node
}

//TODO: Eliminate duplicates from the output
func bstSeq(root *Node) []*list.List {
	if root == nil {
		return []*list.List{list.New()}
	}

	leftLists := bstSeq(root.left)
	rightLists := bstSeq(root.right)
	val := root.val
	var ans []*list.List
	for i := range leftLists {
		for j := range rightLists {
			subAns := list.New()
			subAns.PushBack(val)
			subAns.PushBackList(leftLists[i])
			subAns.PushBackList(rightLists[j])
			ans = append(ans, subAns)

			subAns = list.New()
			subAns.PushBack(val)
			subAns.PushBackList(rightLists[j])
			subAns.PushBackList(leftLists[i])
			ans = append(ans, subAns)
		}
	}
	return ans
}

func createBST(arr []int, lo, hi int) *Node {
	if lo > hi {
		return nil
	}

	mid := lo + (hi-lo)>>1
	val := arr[mid]
	left := createBST(arr, lo, mid-1)
	right := createBST(arr, mid+1, hi)
	return &Node{val: val, left: left, right: right}
}

func printList(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(int))
	}
	fmt.Println()
}

func printLists(lists []*list.List) {
	for _, list := range lists {
		printList(list)
	}
}

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr := []int{1, 2, 3}
	bst := createBST(arr, 0, len(arr)-1)
	lists := bstSeq(bst)
	printLists(lists)
}
