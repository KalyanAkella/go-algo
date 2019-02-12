package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func isOdd(num int) bool {
	return num&1 == 1
}

func GenerateEvenOddLists(node *Node) (*Node, *Node) {
	var even_head, even, odd_head, odd *Node
	for i := node; i != nil; i = i.next {
		if isOdd(i.data) {
			odd_node := &Node{data: i.data}
			if odd != nil {
				odd_node.next = odd
				odd_head = odd_node
				odd = odd_node
			}
			odd_head = odd_node
			odd = odd_node
		} else {
			even_node := &Node{data: i.data}
			if even == nil {
				even_head = even_node
			} else {
				even.next = even_node
			}
			even = even_node
		}
	}
	return even_head, odd_head
}

func printList(node *Node) {
	for i := node; i != nil; i = i.next {
		fmt.Printf("%d ", i.data)
	}
	fmt.Println()
}

func construct(arr []int) *Node {
	var prev, head *Node
	for _, num := range arr {
		new_node := &Node{data: num}
		if prev == nil {
			head = new_node
		} else {
			prev.next = new_node
		}
		prev = new_node
	}
	return head
}

func main() {
	list := construct([]int{0, 1, 2, 3, 4, 5})
	fmt.Println("Given list:")
	printList(list)
	even_list, odd_list := GenerateEvenOddLists(list)
	fmt.Println("Even list:")
	printList(even_list)
	fmt.Println("Odd list:")
	printList(odd_list)
}
