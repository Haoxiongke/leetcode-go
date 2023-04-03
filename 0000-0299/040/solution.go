package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	// 记录是否使用，以免结果重复
	used := make([]bool, len(candidates))
	var backtracking func(index int)

	// 排序方便回溯
	sort.Ints(candidates)

	backtracking = func(index int) {
		if sum(path) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if sum(path) > target {
				break
			}
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			path = append(path, candidates[i])
			used[i] = true
			backtracking(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return result
}

func sum(path []int) (sum int) {
	for _, i := range path {
		sum += i
	}
	return
}
