package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

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
			if i == lens-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}
