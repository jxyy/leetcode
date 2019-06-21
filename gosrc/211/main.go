package main

import "fmt"

type Node struct {
	c          byte
	children   []*Node
	isTerminal bool
}

func NewNode(c byte) *Node {
	return &Node{c, make([]*Node, 26, 26), false}
}

type WordDictionary struct {
	root       *Node
	levelIndex [][][]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	var d = WordDictionary{NewNode(' '), nil}
	return d
}

func (this *WordDictionary) updateLevelIndex(level int, c byte, node *Node) {
	if level < len(this.levelIndex) {
		var index = this.levelIndex[level][c-'a']
		index = append(index, node)
	} else {
		var newIndex = make([][]*Node, 26, 26)
		newIndex[c-'a'] = append(newIndex[c-'a'], node)
		this.levelIndex = append(this.levelIndex, newIndex)
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	var node = this.root
	for i, c := range word {
		var ci = byte(c - 'a')
		if nextNode := node.children[ci]; nextNode != nil {
			node = nextNode
		} else {
			nextNode = NewNode(byte(c))
			node.children[ci] = nextNode
			node = nextNode
			this.updateLevelIndex(i, byte(c), node)
		}
	}
	node.isTerminal = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var pending = []*Node{this.root}
	var isTerminal = false
	for i := range word {
		var c = word[i]
		var nextPending []*Node
		isTerminal = false
		if c == '.' {
			for _, node := range pending {
				for _, nextNode := range node.children {
					if nextNode != nil {
						isTerminal = isTerminal || nextNode.isTerminal
						nextPending = append(nextPending, nextNode)
					}
				}
			}
		} else {
			for _, node := range pending {
				if nextNode := node.children[c-'a']; nextNode != nil {
					isTerminal = isTerminal || nextNode.isTerminal
					nextPending = append(nextPending, nextNode)
				}
			}
		}
		pending = nextPending
		if len(pending) == 0 {
			return false
		}
	}
	return isTerminal
}

func main() {
	var d = Constructor()
	d.AddWord("at")
	d.AddWord("and")
	d.AddWord("an")
	d.AddWord("add")
	// d.AddWord("a")
	// d.AddWord("dad")
	// d.AddWord("mad")
	fmt.Println(d.Search("a"))
	fmt.Println(d.Search(".at"))
	d.AddWord("bat")
	fmt.Println(d.Search(".at"))
	fmt.Println(d.Search("an."))
	fmt.Println(d.Search("a.d."))
	fmt.Println(d.Search("b."))
	fmt.Println(d.Search("a.d"))
	fmt.Println(d.Search("."))
}
