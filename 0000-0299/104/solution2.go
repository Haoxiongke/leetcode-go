package main

import "container/list"

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	depth := 0

	for queue.Len() > 0 {
		lens := queue.Len()
		for i := 0; i < lens; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		depth++
	}
	return depth
}
