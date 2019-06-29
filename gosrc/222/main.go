package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count = 0
	var node = root
	for node != nil {
		count++
		var lh = treeHeight(node.Left)
		var rh = treeHeight(node.Right)
		if lh == rh {
			count += (1 << uint(lh)) - 1
			node = node.Right
		} else {
			// lh > rh
			count += (1 << uint(rh)) - 1
			node = node.Left
		}
	}
	return count
}

func treeHeight(root *TreeNode) int {
	var node = root
	var h = 0
	for node != nil {
		h++
		node = node.Left
	}
	return h
}

func main() {
	var t = &TreeNode{1, nil, nil}
	t.Left = &TreeNode{1, nil, nil}
	t.Right = &TreeNode{1, nil, nil}
	t.Left.Left = &TreeNode{1, nil, nil}
	t.Left.Right = &TreeNode{1, nil, nil}
	t.Right.Left = &TreeNode{1, nil, nil}
	fmt.Println(countNodes(t))
}
