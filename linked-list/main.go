package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(data interface{}) {
	new_node := &Node{data, nil}
	var prev *Node
	for node := ll.head; node != nil; node = node.next {
		prev = node
	}
	if prev == nil {
		ll.head = new_node
	} else {
		prev.next = new_node
	}
}

func (ll *LinkedList) Remove(data interface{}) bool {
	var prev, node *Node
	for node = ll.head; node != nil; node = node.next {
		if node.data == data {
			break
		}
		prev = node
	}
	if node != nil {
		if prev != nil {
			prev.next = node.next
		} else {
			ll.head = node.next
		}
	}
	return node != nil
}

func (ll *LinkedList) Print() {
	fmt.Printf("[")
	for node := ll.head; node != nil; node = node.next {
		fmt.Printf("%v -> ", node.data)
	}
	fmt.Printf("nil]\n")
}

func main() {
	ll := LinkedList{nil}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Print()
	ll.Remove(3)
	ll.Print()
	ll.Remove(1)
	ll.Print()
	ll.Remove(5)
	ll.Print()
}
