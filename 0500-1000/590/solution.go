package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	var order func(node *Node)
	order = func(node *Node) {
		if node == nil {
			return
		}

		if len(node.Children) > 0 {
			for _, child := range node.Children {
				order(child)
			}
		}
		res = append(res, node.Val)
	}

	order(root)
	return res
}
