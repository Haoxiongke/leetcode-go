package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	sort.Ints(candidates)

	var backtracking func(index int)

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
			path = append(path, candidates[i])
			backtracking(i)
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
	return sum
}
