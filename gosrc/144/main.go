package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack = []*TreeNode{root}
	for len(stack) > 0 {
		var l = len(stack)
		var node = stack[l-1]
		stack = stack[:l-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func main() {

}
