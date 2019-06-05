package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
}

func isValidBST(root *TreeNode) bool {
	var min, max int = -0x100000000, 0xffffffff
	return isValid(root, min, max)
}

func main() {
	fmt.Println(isValidBST(&TreeNode{2147483647, nil, nil}))
}
