package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rdfs(node *TreeNode, depth int, res []int) []int {
	if depth == len(res) {
		res = append(res, node.Val)
	}
	if node.Right != nil {
		res = rdfs(node.Right, depth+1, res)
	}
	if node.Left != nil {
		res = rdfs(node.Left, depth+1, res)
	}
	return res
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return rdfs(root, 0, []int{})
}

func main() {
	var t = &TreeNode{1, nil, nil}
	t.Left = &TreeNode{2, nil, nil}
	t.Right = &TreeNode{3, nil, nil}
	t.Left.Right = &TreeNode{5, nil, nil}
	// t.Right.Right = &TreeNode{4, nil, nil}
	fmt.Println(rightSideView(t))
}
