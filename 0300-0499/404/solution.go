package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	var traversal func(node *TreeNode, flag bool)
	traversal = func(node *TreeNode, flag bool) {
		if flag && node.Left == nil && node.Right == nil {
			result += node.Val
		}

		if node.Left != nil {
			traversal(node.Left, true)
		}
		if node.Right != nil {
			traversal(node.Right, false)
		}
	}
	traversal(root, false)
	return result
}
