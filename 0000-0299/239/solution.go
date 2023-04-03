package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {

	if k == 1 {
		return nums
	}

	var queue, result []int

	//queue := make([]int, 0)
	//result := make([]int, 0)
	// 单调队列，求滑动窗口最值。单调递减求最大值
	for i := 0; i < len(nums); i++ {
		for len(queue) != 0 && i-k+1 > queue[0] {
			queue = queue[1:]
		}

		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		if i+1 >= k {
			result = append(result, nums[queue[0]])
		}
	}

	return result
}

func main() {
	window := maxSlidingWindow([]int{7, 2, 4}, 2)
	for _, w := range window {
		fmt.Println(w)
	}
}
