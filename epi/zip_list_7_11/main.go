package main

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

func reverse(head *Node) *Node {
	var result *Node
	for curr := head; curr != nil; {
		next := curr.next
		curr.next = result
		result = curr
		curr = next
	}
	return result
}

func zip(head *Node) {
	fast, slow := head, head
	for slow != nil && slow.next != nil && fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}
	first_part, second_part := head, reverse(slow)
	for first_part != nil && second_part != nil {
		firTemp := first_part.next
		secTemp := second_part.next
		if first_part != second_part {
			first_part.next = second_part
		}
		if second_part != firTemp {
			second_part.next = firTemp
		}
		second_part = secTemp
		first_part = firTemp
	}
}

func printList(head *Node, expCount int) {
	for i := 0; i < expCount && head != nil; i++ {
		fmt.Print(head.val, "->")
		head = head.next
	}
	fmt.Println()
}

func main() {
	head := &Node{"l0", &Node{"l1", &Node{"l2", &Node{"l3", &Node{"l4", &Node{"l5", &Node{"l6", nil}}}}}}}
	printList(head, 10)
	zip(head)
	printList(head, 10)
	head = &Node{"l0", &Node{"l1", &Node{"l2", &Node{"l3", &Node{"l4", &Node{"l5", nil}}}}}}
	printList(head, 10)
	zip(head)
	printList(head, 10)
}
