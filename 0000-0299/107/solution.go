package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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
	reverse(res)
	return res
}

func reverse(res [][]int) {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l, r = l+1, r-1
	}
}
