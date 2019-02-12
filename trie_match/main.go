package main

import "fmt"

type Node struct {
	val         rune
	next_lookup map[rune]*Node
	starred     bool
}

func (n *Node) skipToNextNode() *Node {
	for k, v := range n.next_lookup {
		if k != n.val {
			return v
		}
	}
	panic("Unexpected. Should have found the next node")
}

func constructTrie(input string) *Node {
	var trie_head, prev_node *Node
	trie_head = &Node{val: '^', next_lookup: make(map[rune]*Node)}
	prev_node = trie_head
	for i := 0; i < len(input); i++ {
		c := rune(input[i])
		if c == '*' {
			prev_node.next_lookup[prev_node.val] = prev_node
			prev_node.starred = true
		} else {
			curr_node := &Node{val: c, next_lookup: make(map[rune]*Node)}
			prev_node.next_lookup[c] = curr_node
			prev_node = curr_node
		}
	}
	last_node := &Node{val: '$', next_lookup: make(map[rune]*Node)}
	prev_node.next_lookup['$'] = last_node
	return trie_head
}

func match(input string, trie_head *Node) bool {
	curr_node := trie_head
	for i := 0; i < len(input); i++ {
		c := rune(input[i])
		if next_node, present := curr_node.next_lookup[c]; present {
			fmt.Println("value matched", string(c))
			curr_node = next_node
		} else if dot_node, present := curr_node.next_lookup['.']; present {
			fmt.Println("Dot node matched", string(c))
			curr_node = dot_node
		} else if curr_node.starred {
			fmt.Println("star node matched", string(c))
			i--
			curr_node = curr_node.skipToNextNode()
		} else {
			fmt.Println("did not match", string(c))
			return false
		}
	}
	return true
}

func isMatch(s string, p string) bool {
	trie := constructTrie(p)
	input := fmt.Sprintf("%s$", s)
	return match(input, trie)
}

func main() {
	fmt.Println("vim-go")
}
