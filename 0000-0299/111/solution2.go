package main

import "container/list"

func minDepth2(root *TreeNode) int {
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
			if node.Left == nil && node.Right == nil {
				return depth + 1
			}
		}
		depth++
	}
	return depth
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	println(minDepth2(root))
}
