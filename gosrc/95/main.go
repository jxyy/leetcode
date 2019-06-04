package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cache [][]*TreeNode

func makeTree(tplRoot *TreeNode, addon int) *TreeNode {
	if tplRoot == nil {
		return nil
	}
	var newRoot = &TreeNode{tplRoot.Val + addon, nil, nil}
	var tplStack = make([]*TreeNode, 0, 20)
	var newStack = make([]*TreeNode, 0, 20)
	tplStack = append(tplStack, tplRoot)
	newStack = append(newStack, newRoot)
	for len(tplStack) > 0 {
		var l = len(tplStack)
		var tplNode = tplStack[l-1]
		var newNode = newStack[l-1]
		tplStack = tplStack[:l-1]
		newStack = newStack[:l-1]
		if tplNode.Left != nil {
			newNode.Left = &TreeNode{tplNode.Left.Val + addon, nil, nil}
			tplStack = append(tplStack, tplNode.Left)
			newStack = append(newStack, newNode.Left)
		}
		if tplNode.Right != nil {
			newNode.Right = &TreeNode{tplNode.Right.Val + addon, nil, nil}
			tplStack = append(tplStack, tplNode.Right)
			newStack = append(newStack, newNode.Right)
		}

	}
	return newRoot
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

func gen(n int) []*TreeNode {
	if cache[n] != nil {
		return cache[n]
	}
	var ns []*TreeNode
	for i := 1; i < n+1; i++ {
		var left, right = cache[i-1], cache[n-i]
		if left == nil {
			left = gen(i - 1)
		}
		if len(left) == 0 {
			left = append(left, nil)
		}
		if right == nil {
			right = gen(n - i)
		}
		if len(right) == 0 {
			right = append(right, nil)
		}
		for _, ln := range left {
			for _, rn := range right {
				var lc, rc *TreeNode
				if ln != nil {
					lc = makeTree(ln, 0)
				}
				if rn != nil {
					rc = makeTree(rn, i)
				}
				var root = &TreeNode{i, lc, rc}
				ns = append(ns, root)
			}
		}
	}
	cache[n] = ns
	return cache[n]
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	cache = make([][]*TreeNode, n+1)
	cache[0] = []*TreeNode{}
	cache[1] = []*TreeNode{&TreeNode{1, nil, nil}}
	gen(n)
	return cache[n]
}

func main() {
	ts := generateTrees(0)
	// fmt.Println(len(ts), ts)
	for _, t := range ts {
		printTree(t)
	}
	// printTree(&TreeNode{1, nil, &TreeNode{666, nil, nil}})
}
