package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p {
		return root
	}
	if root == q {
		return root
	}

	var l = lowestCommonAncestor(root.Left, p, q)
	var r = lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	}
	return r
}

func main() {

}
