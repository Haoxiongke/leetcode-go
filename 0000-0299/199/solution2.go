package main

// 递归方法解决，从右子树开始遍历，加入每层第一个达到的元素即可
func rightSideView2(root *TreeNode) []int {
	var res []int
	var order func(node *TreeNode, depth int)

	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		order(node.Right, depth+1)
		order(node.Left, depth+1)
	}
	order(root, 0)
	return res
}
