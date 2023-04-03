package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	result := make([]int, 0)

	for _, num := range nums {
		m[num]++
	}

	for key, _ := range m {
		result = append(result, key)
	}

	sort.Slice(result, func(i, j int) bool {
		return m[result[i]] > m[result[j]]
	})

	return result[:k]
}
