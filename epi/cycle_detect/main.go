package main

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

func cycleNode(head *Node) *Node {
	fast, slow := head, head
	for fast != nil && slow != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			break
		}
	}
	if slow == nil || fast == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

func main() {
	node_a := &Node{val: "a"}
	node_b := &Node{val: "b"}
	node_c := &Node{val: "c"}
	node_d := &Node{val: "d"}
	node_e := &Node{val: "e"}
	node_f := &Node{val: "f"}
	node_g := &Node{val: "g"}

	node_a.next = node_b
	node_b.next = node_c
	node_c.next = node_d
	node_d.next = node_e
	node_e.next = node_f
	node_f.next = node_g
	node_g.next = node_c

	result := cycleNode(node_a)
	fmt.Println(result.val)
}
