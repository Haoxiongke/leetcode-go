package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var preOrder func(node *TreeNode, path string)
	preOrder = func(node *TreeNode, path string) {
		//if node == nil {
		//	return
		//}
		//
		//var prePath string
		//if path == "" {
		//	prePath = strconv.Itoa(node.Val)
		//} else {
		//	prePath = path + "->" + strconv.Itoa(node.Val)
		//}
		//
		//if node.Left == nil && node.Right == nil {
		//	res = append(res, prePath)
		//	return
		//}
		//
		//preOrder(node.Left, prePath)
		//preOrder(node.Right, prePath)
		if node.Left == nil && node.Right == nil {
			path = path + strconv.Itoa(node.Val)
			res = append(res, path)
			return
		}

		path = path + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			preOrder(node.Left, path)
		}

		if node.Right != nil {
			preOrder(node.Right, path)
		}

	}
	preOrder(root, "")
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	paths := binaryTreePaths(root)
	fmt.Println(paths)
}
