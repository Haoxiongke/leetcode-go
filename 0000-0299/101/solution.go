package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var recursion func(left *TreeNode, right *TreeNode) bool
	recursion = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil && right != nil {
			return false
		}
		if left != nil && right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		outside := recursion(left.Left, right.Right)
		inside := recursion(left.Right, right.Left)
		return outside && inside
	}
	return recursion(root.Left, root.Right)
}
