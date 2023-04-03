package main

import "strconv"

func binaryTreePaths2(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}
	var nodeQueue []*TreeNode
	var pathQueue []string
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}
