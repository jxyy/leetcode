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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	var l = len(postorder)
	var v = postorder[l-1]
	var root = &TreeNode{v, nil, nil}
	var i = find(inorder, v)
	root.Left = buildTree(inorder[:i], postorder[:i])
	if i < l-1 {
		root.Right = buildTree(inorder[i+1:l], postorder[i:l-1])
	}
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
	// 	var t = buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	// 	printTree(t)
	fmt.Println([]int{1, 2, 3}[:0])
}
