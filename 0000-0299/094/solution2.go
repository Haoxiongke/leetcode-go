package main

import "container/list"

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	stack := list.New()
	cur := root

	for cur != nil || stack.Len() > 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		cur = node.Right
	}
	return res
}
