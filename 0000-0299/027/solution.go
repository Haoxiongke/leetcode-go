package main

// 二分移动
func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		for l <= r && nums[l] != val {
			l++
		}

		for l <= r && nums[r] == val {
			r--
		}

		for l < r {
			nums[l] = nums[r]
			l++
			r--
		}
	}
	return l
}
