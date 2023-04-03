package main

type Node struct {
	Val      int
	Children []*Node
}

// 多叉树前序遍历
func preorder(root *Node) []int {
	var res []int
	var order func(node *Node)
	order = func(node *Node) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		if len(node.Children) > 0 {
			for _, child := range node.Children {
				order(child)
			}
		}
	}

	order(root)

	return res
}
