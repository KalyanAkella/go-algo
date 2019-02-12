package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func reverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	currNode := reverse(head.next)
	newNode := &Node{val: head.val}

	if currNode == nil {
		return newNode
	} else {
		temp := currNode
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
		return currNode
	}
}

func printList(head *Node) {
	if head == nil {
		return
	}
	fmt.Print(head.val, "->")
	printList(head.next)
}

func main() {
	head := &Node{3, &Node{5, &Node{7, &Node{9, &Node{11, nil}}}}}
	printList(head)
	fmt.Println()
	result := reverse(head)
	printList(result)
	fmt.Println()
}
