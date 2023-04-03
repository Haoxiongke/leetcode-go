package main

import "container/list"

func postorderTraversal2(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)

		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		res = append(res, node.Val)
	}

	reverse(res)

	return res
}

func reverse(res []int) {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l, r = l+1, r-1
	}
}
