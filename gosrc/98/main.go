package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var nextStack []*TreeNode
	var fromLeftToRight = true
	nextStack = append(nextStack, root)
	for len(nextStack) > 0 {
		var currStack = nextStack
		nextStack = []*TreeNode{}
		var nums = []int{}
		for i := len(currStack) - 1; i >= 0; i-- {
			node := currStack[i]
			nums = append(nums, node.Val)
			if fromLeftToRight {
				if node.Left != nil {
					nextStack = append(nextStack, node.Left)
				}
				if node.Right != nil {
					nextStack = append(nextStack, node.Right)
				}
			} else {
				if node.Right != nil {
					nextStack = append(nextStack, node.Right)
				}
				if node.Left != nil {
					nextStack = append(nextStack, node.Left)
				}
			}

		}
		fromLeftToRight = !fromLeftToRight
		res = append(res, nums)
	}
	return res
}

func main() {

}
