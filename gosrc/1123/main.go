package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}
	if node.Left == nil && node.Right == nil {
		return node, depth
	}

	var llca, ld = dfs(node.Left, depth+1)
	var rlca, rd = dfs(node.Right, depth+1)
	if llca != nil && rlca != nil && ld == rd {
		return node, ld
	}
	if ld > rd {
		return llca, ld
	}
	return rlca, rd
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var lca, _ = dfs(root, 0)
	return lca
}

func main() {
	var t = &TreeNode{1, nil, nil}
	t.Left = &TreeNode{2, nil, nil}
	// t.Right = &TreeNode{3, nil, nil}
	// t.Left.Left = &TreeNode{4, nil, nil}
	// t.Left.Right = &TreeNode{5, nil, nil}
	fmt.Println(lcaDeepestLeaves(t))
}
