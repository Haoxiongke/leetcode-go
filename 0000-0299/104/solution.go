package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		return 1 + max(getDepth(node.Left), getDepth(node.Right))
	}
	return getDepth(root)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
