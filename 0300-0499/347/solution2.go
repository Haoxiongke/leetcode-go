package main

import "container/heap"

// 利用优先级队列，小顶堆
type Item struct {
	value int
	count int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for key, count := range m {
		item := Item{key, count}
		if h.Len() < k {
			heap.Push(h, item)
		} else if h.Len() > 0 && count > (*h)[0].count {
			heap.Pop(h)
			heap.Push(h, item)
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = (*h)[i].value
	}

	return result
}
