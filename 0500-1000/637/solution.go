package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res [][]int

	var order func(node *TreeNode, depth int)

	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		order(node.Left, depth+1)
		order(node.Right, depth+1)
		res[depth] = append(res[depth], node.Val)
	}

	order(root, 0)

	var resFloat []float64
	for _, list := range res {
		resFloat = append(resFloat, avg(list))
	}
	return resFloat
}

func avg(list []int) float64 {
	var sum int
	for _, data := range list {
		sum += data
	}
	return float64(sum) / float64(len(list))
}
