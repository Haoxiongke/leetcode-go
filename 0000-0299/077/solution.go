package main

func combine(n int, k int) [][]int {
	var result [][]int
	var path []int

	var backtracking func(n, k, startIndex int)

	backtracking = func(n, k, startIndex int) {
		if len(path) == k {
			// 切片需要拷贝，要不然会被覆盖
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := startIndex; i <= n; i++ {
			// 剪枝优化
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)
	return result
}
