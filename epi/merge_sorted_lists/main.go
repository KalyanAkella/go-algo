package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func merge(head1, head2 *Node) *Node {
	var resultHead, resultNode *Node
	for head1 != nil && head2 != nil {
		var newNode *Node
		if head1.val < head2.val {
			newNode = &Node{val: head1.val}
			head1 = head1.next
		} else if head1.val > head2.val {
			newNode = &Node{val: head2.val}
			head2 = head2.next
		} else {
			newNode = &Node{val: head1.val}
			head1 = head1.next
			head2 = head2.next
		}
		if resultNode == nil {
			resultNode = newNode
			resultHead = newNode
		} else {
			resultNode.next = newNode
			resultNode = newNode
		}
	}

	var remaining *Node
	if head1 != nil {
		remaining = head1
	} else {
		remaining = head2
	}

	for remaining != nil {
		newNode := &Node{val: remaining.val}
		remaining = remaining.next
		if resultNode == nil {
			resultNode = newNode
			resultHead = newNode
		} else {
			resultNode.next = newNode
			resultNode = newNode
		}
	}
	return resultHead
}

func printList(head *Node) {
	if head == nil {
		return
	}
	fmt.Print(head.val, "->")
	printList(head.next)
}

func main() {
	head1 := &Node{2, &Node{5, &Node{7, nil}}}
	head2 := &Node{3, &Node{11, nil}}
	result := merge(head1, head2)
	printList(result)
}
