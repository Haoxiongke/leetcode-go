package main

import "container/list"

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	return root
}
