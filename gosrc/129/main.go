package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, num int, sum *int) {
	num *= 10
	num += node.Val
	if node.Left == nil && node.Right == nil {
		*sum += num
		return
	}
	if node.Left != nil {
		dfs(node.Left, num, sum)
	}
	if node.Right != nil {
		dfs(node.Right, num, sum)
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum = 0
	dfs(root, 0, &sum)
	return sum
}

func main() {
	var tree = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(sumNumbers(tree))
}
