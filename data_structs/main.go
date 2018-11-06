package main

import (
	ll "algo/data_structs/linked_list"
	"fmt"
)

func testRemoveDups() {
	head := &ll.Node{Value: 10}
	head.Next = &ll.Node{Value: 8}
	head.Next.Next = &ll.Node{Value: 9}
	head.Next.Next.Next = &ll.Node{Value: 8} //	duplicate
	head.Next.Next.Next.Next = &ll.Node{Value: 3}

	ll.Print(head)
	ll.RemoveDuplicates(head)
	ll.Print(head)
}

func testReturnKthLastElement() {
	head := &ll.Node{Value: 10}
	head.Next = &ll.Node{Value: 9}
	head.Next.Next = &ll.Node{Value: 8}
	head.Next.Next.Next = &ll.Node{Value: 7}
	head.Next.Next.Next.Next = &ll.Node{Value: 6}

	fmt.Println(ll.ReturnFromLast(head, 1))
	fmt.Println(ll.ReturnFromLast(head, 2))
	fmt.Println(ll.ReturnFromLast(head, 3))
	fmt.Println(ll.ReturnFromLast(head, 4))
	fmt.Println(ll.ReturnFromLast(head, 5))
	fmt.Println(ll.ReturnFromLast(head, 6))
	fmt.Println(ll.ReturnFromLast(head, 7))
}

func testDeleteMiddleNode() {
	head := &ll.Node{Value: "a"}
	head.Next = &ll.Node{Value: "b"}
	head.Next.Next = &ll.Node{Value: "c"}
	head.Next.Next.Next = &ll.Node{Value: "d"}
	head.Next.Next.Next.Next = &ll.Node{Value: "e"}
	head.Next.Next.Next.Next.Next = &ll.Node{Value: "f"}

	ll.Print(head)
	fmt.Println(ll.DeleteMiddleNode(head))
	ll.Print(head)
}

func testPartition() {
	head := &ll.Node{Value: 3}
	head.Next = &ll.Node{Value: 5}
	head.Next.Next = &ll.Node{Value: 8}
	head.Next.Next.Next = &ll.Node{Value: 5}
	head.Next.Next.Next.Next = &ll.Node{Value: 10}
	head.Next.Next.Next.Next.Next = &ll.Node{Value: 2}
	head.Next.Next.Next.Next.Next.Next = &ll.Node{Value: 1}

	ll.Print(head)
	ll.Print(ll.Partition(head, 5))
}

func testSumDigitsInReverse() {
	head1 := &ll.Node{Value: 7, Next: &ll.Node{Value: 1, Next: &ll.Node{Value: 6}}}
	head2 := &ll.Node{Value: 5, Next: &ll.Node{Value: 9, Next: &ll.Node{Value: 2}}}

	ll.Print(ll.SumListsWithDigitsInReverse(head1, head2))
}

func testSumDigitsInForward() {
	head1 := &ll.Node{Value: 6, Next: &ll.Node{Value: 1, Next: &ll.Node{Value: 7}}}
	head2 := &ll.Node{Value: 2, Next: &ll.Node{Value: 9, Next: &ll.Node{Value: 5}}}

	ll.Print(ll.SumListsWithDigitsInForward(head1, head2))
}

func testReverse() {
	head := &ll.Node{Value: "a"}
	head.Next = &ll.Node{Value: "b"}
	head.Next.Next = &ll.Node{Value: "c"}
	head.Next.Next.Next = &ll.Node{Value: "d"}
	head.Next.Next.Next.Next = &ll.Node{Value: "e"}
	head.Next.Next.Next.Next.Next = &ll.Node{Value: "f"}

	ll.Print(head)
	ll.Print(ll.Reverse(head))
}

func testPalindrome() {
	head := &ll.Node{Value: 1, Next: &ll.Node{Value: 2, Next: &ll.Node{Value: 2, Next: &ll.Node{Value: 1}}}}

	ll.Print(head)
	fmt.Println(ll.IsPalindrome(head))

	head = &ll.Node{Value: 1, Next: &ll.Node{Value: 2, Next: &ll.Node{Value: 3, Next: &ll.Node{Value: 2, Next: &ll.Node{Value: 1}}}}}

	ll.Print(head)
	fmt.Println(ll.IsPalindrome(head))
}

func testIntersectingNode() {
	head1 := &ll.Node{Value: 6, Next: &ll.Node{Value: 1, Next: &ll.Node{Value: 7, Next: &ll.Node{Value: 4}}}}
	head2 := &ll.Node{Value: 2, Next: &ll.Node{Value: 9, Next: &ll.Node{Value: 5}}}
	fmt.Println(ll.IntersectingNode(head1, head2))

	commonNode := &ll.Node{Value: "hello"}
	commonNode.Next = head1.Next.Next
	head1.Next.Next = commonNode
	ll.Print(head1)

	commonNode.Next = head2.Next
	head2.Next = commonNode
	ll.Print(head2)

	fmt.Println(ll.IntersectingNode(head1, head2))
}

func testDetectCycle() {
	head := &ll.Node{Value: "a"}
	head.Next = &ll.Node{Value: "b"}
	head.Next.Next = &ll.Node{Value: "c"}
	head.Next.Next.Next = &ll.Node{Value: "d"}
	head.Next.Next.Next.Next = &ll.Node{Value: "e"}

	fmt.Println(ll.DetectCycle(head))

	head.Next.Next.Next.Next.Next = head.Next.Next
	fmt.Println(ll.DetectCycle(head))
}

func main() {
	testDetectCycle()
}
