package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var c = 0
	var kthNode *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}
		c++
		if c == k {
			kthNode = node
			return
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return kthNode.Val
}

func main() {
	var r = &TreeNode{3, nil, nil}
	r.Left = &TreeNode{1, nil, nil}
	r.Left.Right = &TreeNode{2, nil, nil}
	r.Right = &TreeNode{4, nil, nil}
	fmt.Println(kthSmallest(r, 1))
	fmt.Println(kthSmallest(r, 2))
	fmt.Println(kthSmallest(r, 3))
	fmt.Println(kthSmallest(r, 4))
}
