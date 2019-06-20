package main

import "fmt"

type Trie struct {
	root *Node
}

type Node struct {
	b          byte
	isTerminal bool
	children   map[byte]*Node
}

func NewNode(b byte) *Node {
	var node = Node{b, false, nil}
	node.children = make(map[byte]*Node)
	return &node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var t = Trie{NewNode(' ')}
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var node = this.root
	for i := range word {
		var c = word[i]
		if nextNode, ok := node.children[c]; ok {
			node = nextNode
		} else {
			nextNode = NewNode(c)
			node.children[c] = nextNode
			node = nextNode
		}
	}
	node.isTerminal = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var node = this.root
	for i := range word {
		var c = word[i]
		if nextNode, ok := node.children[c]; ok {
			node = nextNode
		} else {
			return false
		}
	}
	return node.isTerminal
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var node = this.root
	for i := range prefix {
		var c = prefix[i]
		if nextNode, ok := node.children[c]; ok {
			node = nextNode
		} else {
			return false
		}
	}
	return true
}

func main() {
	var trie = Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // returns true
	fmt.Println(trie.Search("app"))     // returns false
	fmt.Println(trie.StartsWith("app")) // returns true
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
