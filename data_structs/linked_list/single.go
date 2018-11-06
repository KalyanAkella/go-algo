package linked_list

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func IntersectingNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	tail1, size1 := TailNodeAndSize(head1)
	tail2, size2 := TailNodeAndSize(head2)

	if tail1 != tail2 {
		return nil
	}

	var short, long *Node
	if size1 < size2 {
		short = head1
		long = head2
	} else {
		short = head2
		long = head1
	}

	p1 := GetNthNode(long, abs(size1-size2))
	p2 := short

	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

/*
O(N) space complexity, O(N) time complexity where N is the number of nodes in LinkedList up until the looping node
*/
func DetectCycleWithLookup(head *Node) *Node {
	lookup := make(map[*Node]struct{})
	for temp := head; temp != nil; temp = temp.Next {
		if _, present := lookup[temp]; present {
			return temp
		} else {
			lookup[temp] = struct{}{}
		}
	}
	return nil
}

func DetectCycle(head *Node) *Node {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			break // collision, but may not be the looping node
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

func GetNthNode(head *Node, n int) *Node {
	for temp := head; temp != nil; temp = temp.Next {
		n--
		if n < 0 {
			return temp
		}
	}
	return nil
}

func TailNodeAndSize(head *Node) (*Node, int) {
	var prev *Node
	count := 0
	for temp := head; temp != nil; temp = temp.Next {
		prev = temp
		count++
	}
	return prev, count
}

func Print(head *Node) {
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Printf("%v -> ", temp.Value)
	}
	fmt.Println("END")
}

func IsPalindrome(head *Node) bool {
	slow, fast := head, head
	stk := list.New()
	for slow != nil && fast != nil {
		stk.PushBack(slow.Value)
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			stk.Remove(stk.Back())
		}
	}

	for slow != nil && stk.Len() > 0 {
		value := stk.Remove(stk.Back())
		if value != slow.Value {
			return false
		}
		slow = slow.Next
	}

	return slow == nil && stk.Len() == 0
}

func Reverse(head *Node) *Node {
	var resultHead *Node
	for temp := head; temp != nil; temp = temp.Next {
		newNode := &Node{Value: temp.Value}
		if resultHead == nil {
			resultHead = newNode
		} else {
			newNode.Next = resultHead
			resultHead = newNode
		}
	}
	return resultHead
}

func SumListsWithDigitsInForward(head1, head2 *Node) *Node {
	return Reverse(SumListsWithDigitsInReverse(Reverse(head1), Reverse(head2)))
}

func SumListsWithDigitsInReverse(head1, head2 *Node) *Node {
	var result, resultHead *Node
	carry := 0
	for head1 != nil && head2 != nil {
		sum := head1.Value.(int) + head2.Value.(int) + carry
		newNode := &Node{Value: sum % 10}
		if resultHead == nil {
			resultHead = newNode
			result = newNode
		} else {
			result.Next = newNode
			result = newNode
		}
		carry = sum / 10
		head1 = head1.Next
		head2 = head2.Next
	}
	for head1 != nil {
		sum := head1.Value.(int) + carry
		newNode := &Node{Value: sum % 10}
		result.Next = newNode // assume result ptr is not nil
		result = newNode
		carry = sum / 10
		head1 = head1.Next
	}
	for head2 != nil {
		sum := head2.Value.(int) + carry
		newNode := &Node{Value: sum % 10}
		result.Next = newNode // assume result ptr is not nil
		result = newNode
		carry = sum / 10
		head2 = head2.Next
	}
	if carry > 0 {
		newNode := &Node{Value: carry}
		result.Next = newNode
	}
	return resultHead
}

func Partition(head *Node, k int) *Node {
	var lessNodes, lessNodesHead *Node
	var greaterOrEqualNodes, greaterOrEqualNodesHead *Node
	for temp := head; temp != nil; temp = temp.Next {
		newNode := &Node{Value: temp.Value}
		if temp.Value.(int) < k {
			if lessNodes == nil {
				lessNodes = newNode
				lessNodesHead = lessNodes
			} else {
				lessNodes.Next = newNode
				lessNodes = newNode
			}
		} else {
			if greaterOrEqualNodes == nil {
				greaterOrEqualNodes = newNode
				greaterOrEqualNodesHead = greaterOrEqualNodes
			} else {
				greaterOrEqualNodes.Next = newNode
				greaterOrEqualNodes = newNode
			}
		}
	}
	lessNodes.Next = greaterOrEqualNodesHead
	return lessNodesHead
}

func DeleteMiddleNode(head *Node) *Node {
	var prev *Node
	slow, fast := head, head
	for slow != nil && fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	prev.Next = slow.Next
	slow.Next = nil
	return slow
}

func ReturnFromLast(head *Node, k int) *Node {
	p1 := head
	for i := 0; p1 != nil && i < k; i++ {
		p1 = p1.Next
	}

	if p1 == nil {
		return nil
	}

	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

func RemoveDuplicates(head *Node) {
	for i := head; i != nil; i = i.Next {
		prev := i
		for j := i.Next; j != nil; j = j.Next {
			if i.Value == j.Value {
				prev.Next = j.Next
				j = prev
			} else {
				prev = j
			}
		}
	}
}
