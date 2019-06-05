package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(l []int, n int) int {
	for i, ln := range l {
		if ln == n {
			return i
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var root = &TreeNode{preorder[0], nil, nil}
	var i = find(inorder, root.Val)
	if i > 0 {
		root.Left = buildTree(preorder[1:1+i], inorder[:i])
	}
	root.Right = buildTree(preorder[1+i:], inorder[1+i:])
	return root
}

func printTree(root *TreeNode) {
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			fmt.Print(node.Val, ", ")
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		} else {
			fmt.Print("null, ")
		}
	}
	fmt.Println("")
}

func main() {
	var t = buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printTree(t)
}
