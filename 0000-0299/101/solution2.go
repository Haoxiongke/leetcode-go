package main

import "container/list"

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}
