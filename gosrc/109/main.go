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

func buildTree(head *ListNode, n int) *TreeNode {
	if n == 0 {
		return nil
	}
	var node = head
	var root *TreeNode
	for i := 0; i < n/2; i++ {
		node = node.Next
	}
	root = &TreeNode{node.Val, nil, nil}
	root.Left = buildTree(head, n/2)
	root.Right = buildTree(node.Next, n-n/2-1)
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	var node = head
	var n = 0
	for node != nil {
		n++
		node = node.Next
	}
	return buildTree(head, n)
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
