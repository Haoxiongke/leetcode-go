package main

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var path []int

	var backtracking func(k, n, startIndex int)

	backtracking = func(k, n, startIndex int) {
		if len(path) == k && sum(path) == n {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
		}

		for i := startIndex; i <= 9; i++ {
			if 9-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backtracking(k, n, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(k, n, 1)
	return result
}

func sum(path []int) (sum int) {
	for _, num := range path {
		sum += num
	}
	return
}
