package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// fmt.Println("root", root, root.Left, root.Right)
	var left, right = root.Left, root.Right
	if left != nil {
		flatten(left)
		root.Right = left
		fmt.Println("set root right:", root, left)
		root.Left = nil
		var node = root
		for node.Right != nil {
			node = node.Right
		}
		// fmt.Println("node", node)
		flatten(right)
		node.Right = right
		fmt.Println("set node right:", node, right)
	} else {
		flatten(right)
	}

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
	var t = &TreeNode{1, nil, nil}
	t.Left = &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}
	t.Right = &TreeNode{5, nil, &TreeNode{6, nil, nil}}

	// var t = &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	flatten(t)

	var tt = t
	for tt != nil {
		fmt.Print(tt.Val, " -> ")
		tt = tt.Right
	}
	fmt.Println("")
}
