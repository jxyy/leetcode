package main

import "fmt"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	pending []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var iter = BSTIterator{nil}
	iter.leftSearch(root)
	return iter
}

func (this *BSTIterator) leftSearch(node *TreeNode) {
	if node == nil {
		return
	}
	var n = node
	for n != nil {
		this.pending = append(this.pending, n)
		n = n.Left
	}
}

/*@return the next smallest number */
func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return 0
	}
	var l = len(this.pending)
	var nextNode = this.pending[l-1]
	this.pending = this.pending[:l-1]
	this.leftSearch(nextNode.Right)
	return nextNode.Val
}

/*@return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.pending != nil && len(this.pending) > 0
}

/**
Your BSTIterator object will be instantiated and called as such:
obj := Constructor(root);
param_1 := obj.Next();
param_2 := obj.HasNext();
*/

func main() {
	var root = &TreeNode{7, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Right.Left = &TreeNode{9, nil, nil}
	root.Right.Right = &TreeNode{20, nil, nil}
	var iter = Constructor(root)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
