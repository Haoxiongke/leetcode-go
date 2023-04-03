package main

import "container/list"

func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	depth := 0

	for queue.Len() > 0 {
		res = append(res, []int{})
		lens := queue.Len()
		for i := 0; i < lens; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			res[depth] = append(res[depth], node.Val)
		}
		depth++
	}
	return res
}
