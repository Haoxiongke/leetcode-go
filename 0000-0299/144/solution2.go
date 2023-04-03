package main

import "container/list"

func preorderTraversal2(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		res = append(res, node.Val)
	}

	return res
}
