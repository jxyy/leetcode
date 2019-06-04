package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	if root.Left != nil {
		ret = append(ret, inorderTraversal(root.Left)...)
	}
	ret = append(ret, root.Val)
	if root.Right != nil {
		ret = append(ret, inorderTraversal(root.Right)...)
	}
	return ret
}

func main() {

}
