package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var result bool

	var traversal func(node *TreeNode, sum int)
	traversal = func(node *TreeNode, sum int) {
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				result = true || result
			}
			result = false || result
		}

		if node.Left != nil {
			traversal(node.Left, sum)
		}

		if node.Right != nil {
			traversal(node.Right, sum)
		}
	}

	traversal(root, 0)
	return result
}
