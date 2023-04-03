package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if getHeight(root) == -1 {
		return false
	}
	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l, r := getHeight(node.Left), getHeight(node.Right)
	if l == -1 || r == -1 {
		return -1
	}

	if l-r > 1 || r-l > 1 {
		return -1
	}

	return 1 + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
