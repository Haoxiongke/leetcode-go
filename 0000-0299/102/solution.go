package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	var order func(node *TreeNode, depth int)
	depth := 0

	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)
		order(node.Left, depth+1)
		order(node.Right, depth+1)
	}

	order(root, depth)
	return res
}
