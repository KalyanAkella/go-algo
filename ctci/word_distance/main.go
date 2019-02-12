package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	ch     byte
	next   map[byte]*Node
	offset int
}

func newNode(ch byte) *Node {
	return &Node{ch: ch, next: make(map[byte]*Node)}
}

type Trie struct {
	root *Node
}

func (this *Trie) Insert(word string, offset int) {
	currNode := this.root
	for i := range word {
		c := word[i]
		if next, present := currNode.next[c]; present {
			currNode = next
		} else {
			currNode.next[c] = newNode(c)
			currNode = currNode.next[c]
		}
	}
	currNode.offset = offset
}

func (this *Trie) Search(word string) int {
	currNode := this.root
	for i := range word {
		c := word[i]
		if next, present := currNode.next[c]; present {
			currNode = next
		} else {
			return 0
		}
	}
	return currNode.offset
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func buildTrie(scanner *bufio.Scanner) *Trie {
	trie := &Trie{root: newNode('^')}
	for offset := 1; scanner.Scan(); offset++ {
		word := scanner.Text()
		trie.Insert(word, offset)
	}
	return trie
}

func wordDistance(fileName string, word1, word2 string) int {
	fd, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(bufio.NewReader(fd))
	scanner.Split(bufio.ScanLines)

	trie := buildTrie(scanner)
	return abs(trie.Search(word1) - trie.Search(word2))
}

func main() {
	trie := &Trie{root: newNode('^')}
	trie.Insert("ball", 3)
	trie.Insert("bald", 7)
	fmt.Println(trie.Search("ball"))
	fmt.Println(trie.Search("bald"))
}
